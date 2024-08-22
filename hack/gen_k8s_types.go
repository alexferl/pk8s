package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"slices"
	"sort"
	"strings"
	"text/template"
	"unicode"

	"github.com/ettle/strcase"
	"github.com/pb33f/libopenapi"
	highbase "github.com/pb33f/libopenapi/datamodel/high/base"
	"github.com/pb33f/libopenapi/datamodel/high/v2"
	"github.com/pb33f/libopenapi/orderedmap"
)

type Definition struct {
	XKubernetesGroupVersionKind []struct {
		Group   string `json:"group"`
		Version string `json:"version"`
		Kind    string `json:"kind"`
	} `json:"x-kubernetes-group-version-kind"`
}

type Swagger struct {
	Definitions map[string]Definition `json:"definitions"`
}

type File struct {
	Name    string
	Package string
	Imports []string
	Struct  Struct
}

type Struct struct {
	Name        string
	Description string
	Fields      []Field
}

type Field struct {
	Comment string
	Name    string
	Type    string
	Tag     string
}

const structTemplate = `// Code generated; DO NOT EDIT.

package {{.Package}}
{{ if .Imports}}
import (
{{- range .Imports}}
	"{{.}}"
{{- end}}
){{end}}

// {{.Struct.Name}} {{.Struct.Description}}
type {{.Struct.Name}} struct {
{{- range .Struct.Fields}}
{{- if .Comment}}
	// {{.Comment}}
{{- end}}
	{{.Name}} {{.Type}} {{.Tag}}
{{- end}}
}
`

var version *string

func main() {
	version = flag.String("version", "", "Kubernetes version")
	flag.Parse()

	url := fmt.Sprintf("https://raw.githubusercontent.com/kubernetes/kubernetes/%s/api/openapi-spec/swagger.json", *version)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("failed to read response body: %v", err)
	}

	p := New()
	p.Process(data)

	tmpl, err := template.New("structTemplate").Parse(structTemplate)
	if err != nil {
		log.Fatalf("failed creating template: %v", err)
	}

	if err := os.MkdirAll("k8s", 0o755); err != nil {
		log.Fatalf("failed to create k8s directory: %v", err)
	}

	for _, file := range p.files {
		out, err := os.Create(fmt.Sprintf("k8s/%s", file.Name))
		if err != nil {
			log.Fatalf("failed creating file: %v", err)
		}
		defer out.Close()

		err = tmpl.Execute(out, file)
		if err != nil {
			log.Fatalf("failed executing template: %v", err)
		}
	}

	var swagger Swagger
	err = json.Unmarshal(data, &swagger)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %v", err)
	}

	switchFunction, testFunction := generateSwitchAndTestFunctions(swagger)

	err = os.WriteFile("k8s/api_version_and_kind.go", []byte(switchFunction), 0o644)
	if err != nil {
		log.Fatalf("failed to write to file: %v", err)
	}

	err = os.WriteFile("k8s/api_version_and_kind_test.go", []byte(testFunction), 0o644)
	if err != nil {
		log.Fatalf("failed to write test file: %v", err)
	}

	generateUtils()

	err = runGoFmt("k8s")
	if err != nil {
		log.Fatalf("%v", err)
	}
}

type Processor struct {
	definitions *orderedmap.Map[string, *highbase.SchemaProxy]
	files       map[string]*File
}

func New() *Processor {
	return &Processor{
		files: make(map[string]*File),
	}
}

func (p *Processor) Process(data []byte) {
	document, err := libopenapi.NewDocument(data)
	if err != nil {
		log.Fatalf("failed creating new document: %e", err)
	}

	var errors []error
	var v2Model *libopenapi.DocumentModel[v2.Swagger]

	v2Model, errors = document.BuildV2Model()
	if len(errors) > 0 {
		for i := range errors {
			fmt.Printf("error: %e\n", errors[i])
		}
		log.Fatalf("cannot create v2 model from "+
			"document: %d errors reported", len(errors))
	}

	p.definitions = v2Model.Model.Definitions.Definitions

	for el := p.definitions.First(); el != nil; el = el.Next() {
		p.process(el.Value())
	}
}

func (p *Processor) append(file *File) {
	if len(file.Struct.Fields) < 1 {
		return
	}

	if _, ok := p.files[file.Name]; !ok {
		p.files[file.Name] = file
	}
}

func (p *Processor) process(def *highbase.SchemaProxy) {
	resourceType, versionTag := splitRef(def.GetSchemaKeyNode().Value)
	isJSONSchemaProps := resourceType == "JSONSchemaProps"

	file := &File{
		Name:    fmt.Sprintf("%s_%s.go", strcase.ToSnake(resourceType), strings.ToLower(versionTag)),
		Package: "k8s",
		Struct: Struct{
			Name:        fmt.Sprintf("%s%s", resourceType, versionTag),
			Description: strings.ReplaceAll(def.Schema().Description, "\n", " "),
		},
	}
	defer p.append(file)

	requiredField := def.Schema().Required
	for el := def.Schema().Properties.First(); el != nil; el = el.Next() {
		schema := el.Value().Schema()

		tag := fmt.Sprintf("`json:\"%s", el.Key())
		if !slices.Contains(requiredField, el.Key()) {
			tag += ",omitempty\"`"
		} else {
			tag += "\"`"
		}

		var typ string
		if !slices.Contains(requiredField, el.Key()) {
			typ += "*"
		}

		key := el.Key()

		if !el.Value().IsReference() {
			if isJSONSchemaProps {
				switch el.Key() {
				case "enum":
					typ = "[]map[string]interface{}"
				case "allOf", "anyOf", "oneOf":
					typ = "[]*JSONSchemaPropsV1"
				case "properties", "patternProperties", "definitions":
					typ = "map[string]*JSONSchemaPropsV1"
				default:
					if strings.HasPrefix(key, "$") {
						key = strings.TrimPrefix(key, "$")
					}
				}
			}
			if schema.Format == "" {
				switch schema.Type[0] {
				case "array":
					if isJSONSchemaProps && key == "enum" {
						continue
					}

					if schema.Items.A.Schema().Format != "" {
						typ = fmt.Sprintf("[]%s", schema.Items.A.Schema().Format)
					} else if len(schema.Items.A.Schema().Type) > 0 && schema.Items.A.Schema().Type[0] != "" {
						typ = fmt.Sprintf("[]%s", schema.Items.A.Schema().Type[0])
					}
					if schema.Items.A.IsReference() {
						ref := strings.TrimPrefix(schema.Items.A.GetReference(), "#/definitions/")
						if ref == "io.k8s.apimachinery.pkg.util.intstr.IntOrString" {
							typ += "string"
						} else {
							resourceType, versionTag = splitRef(ref)
							typ = fmt.Sprintf("[]%s%s", resourceType, capitalizeVersion(versionTag))
						}
					}
				case "boolean":
					typ += "bool"
				case "object":
					if typ != "map[string]*JSONSchemaPropsV1" {
						typ += "map[string]string"
					}
				default:
					typ += schema.Type[0]
				}
			} else {
				switch schema.Format {
				case "double":
					typ += "float64"
				default:
					typ += schema.Format
				}
			}
		} else {
			if isJSONSchemaProps {
				switch el.Key() {
				case "additionalItems", "additionalProperties", "default", "example", "items":
					typ = "map[string]interface{}"
				case "externalDocs", "not":
					ref := strings.TrimPrefix(el.Value().GetReference(), "#/definitions/")
					resourceType, versionTag = splitRef(ref)
					typ += fmt.Sprintf("%s%s", resourceType, versionTag)
				}
			} else {
				ref := strings.TrimPrefix(el.Value().GetReference(), "#/definitions/")
				if ref == "io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime" || ref == "io.k8s.apimachinery.pkg.apis.meta.v1.Time" {
					typ += "time.Time"
					if !slices.Contains(file.Imports, "time") {
						file.Imports = append(file.Imports, "time")
					}
				} else if ref == "io.k8s.apimachinery.pkg.util.intstr.IntOrString" || ref == "io.k8s.apimachinery.pkg.api.resource.Quantity" {
					typ += "string"
				} else {
					resourceType, versionTag = splitRef(ref)
					typ += fmt.Sprintf("%s%s", resourceType, versionTag)
				}
			}
		}

		// TODO: handle properly above
		if typ == "RawExtensionRuntime" || typ == "*RawExtensionRuntime" || typ == "*FieldsV1V1" || typ == "StorageVersionSpecV1alpha1" || typ == "*CustomResourceSubresourceStatusV1" {
			typ = "map[string]interface{}"
		}

		field := Field{
			Comment: strings.ReplaceAll(schema.Description, "\n", " "),
			Name:    strcase.ToGoPascal(key),
			Tag:     tag,
			Type:    typ,
		}

		file.Struct.Fields = append(file.Struct.Fields, field)

		if el.Value().IsReference() {
			ref := strings.TrimPrefix(el.Value().GetReference(), "#/definitions/")
			if ref != "io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.JSONSchemaProps" {
				def, _ = p.definitions.Get(ref)
				p.process(def)
			}
		}
	}
}

func splitRef(ref string) (string, string) {
	parts := strings.Split(ref, ".")
	versionTag := parts[len(parts)-2]
	resourceType := parts[len(parts)-1]
	return resourceType, capitalizeVersion(versionTag)
}

func capitalizeVersion(version string) string {
	if len(version) > 0 {
		runes := []rune(version)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
	return version
}

func runGoFmt(dir string) error {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("go fmt failed: %v\n%s", err, output)
	}
	return nil
}

func generateUtils() {
	content := `// Code generated; DO NOT EDIT.

package k8s

import (
	"time"
)

// Bool is a helper routine that returns a pointer to given boolean value.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that returns a pointer to given integer value.
func Int(v int) *int { return &v }

// Int32 is a helper routine that returns a pointer to given integer value.
func Int32(v int32) *int32 { return &v }

// Int64 is a helper routine that returns a pointer to given integer value.
func Int64(v int64) *int64 { return &v }

// Float32 is a helper routine that returns a pointer to given float value.
func Float32(v float32) *float32 { return &v }

// Float64 is a helper routine that returns a pointer to given float value.
func Float64(v float64) *float64 { return &v }

// String is a helper routine that returns a pointer to given string value.
func String(v string) *string { return &v }

// Time is helper routine that returns a pointer to given Time value.
func Time(v time.Time) *time.Time { return &v }
`

	err := os.WriteFile("k8s/utils.go", []byte(content), 0o644)
	if err != nil {
		log.Fatalf("failed writing to file: %v\n", err)
	}
}

func generateSwitchAndTestFunctions(swagger Swagger) (string, string) {
	var cases []string
	var testCases []string
	var testAssertions []string

	for _, def := range swagger.Definitions {
		for _, gvk := range def.XKubernetesGroupVersionKind {
			// Skip
			if gvk.Kind == "DeleteOptions" || gvk.Kind == "WatchEvent" || gvk.Kind == "TokenRequest" || strings.HasSuffix(gvk.Kind, "List") {
				continue
			}

			// Special handling for Event: only include if it's "events.k8s.io/v1"
			if gvk.Kind == "Event" && !(gvk.Group == "events.k8s.io" && gvk.Version == "v1") {
				continue
			}

			// Special handling for Status: skip if it's "resource.k8s.io/v1alpha2"
			if gvk.Kind == "Status" && gvk.Group == "resource.k8s.io" && gvk.Version == "v1alpha2" {
				continue
			}

			var apiVersion string
			if gvk.Group == "" {
				apiVersion = gvk.Version
			} else {
				apiVersion = fmt.Sprintf("%s/%s", gvk.Group, gvk.Version)
			}

			capitalizedVersion := capitalizeVersion(gvk.Version)

			caseStr := fmt.Sprintf("case *%s%s:\n\t\tv.APIVersion = String(\"%s\")\n\t\tv.Kind = String(\"%s\")\n", gvk.Kind, capitalizedVersion, apiVersion, gvk.Kind)
			cases = append(cases, caseStr)

			testCaseStr := fmt.Sprintf(`{
				name:       "%s%s",
				input:      &%s%s{},
				wantApiVer: "%s",
				wantKind:   "%s",
			},`, gvk.Kind, capitalizedVersion, gvk.Kind, capitalizedVersion, apiVersion, gvk.Kind)
			testCases = append(testCases, testCaseStr)

			testAssertionStr := fmt.Sprintf(`case *%s%s:
				assert.Equal(t, tt.wantApiVer, *obj.APIVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
`, gvk.Kind, capitalizedVersion)
			testAssertions = append(testAssertions, testAssertionStr)
		}
	}

	// Sort the cases by kind
	sort.Strings(cases)
	sort.Strings(testCases)
	sort.Strings(testAssertions)

	switchFunction := "// Code generated; DO NOT EDIT.\n\npackage k8s\n\nfunc SetAPIVersionAndKind(i any) {\n\tswitch v := i.(type) {\n"
	for _, caseStr := range cases {
		switchFunction += "\t" + caseStr
	}
	switchFunction += "\t}\n}\n"

	testFunction := `// Code generated; DO NOT EDIT.

package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAPIVersionAndKind(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		wantApiVer string
		wantKind   string
	}{
`
	for _, testCaseStr := range testCases {
		testFunction += "\t\t" + testCaseStr + "\n"
	}
	testFunction += `	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetAPIVersionAndKind(tt.input)

			switch obj := tt.input.(type) {
`
	for _, testAssertionStr := range testAssertions {
		testFunction += "\t\t\t" + testAssertionStr
	}
	testFunction += `			}
		})
	}
}
`

	return switchFunction, testFunction
}
