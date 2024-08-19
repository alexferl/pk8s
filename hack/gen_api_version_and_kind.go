package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"unicode"
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

func capitalizeVersion(version string) string {
	if len(version) > 0 {
		runes := []rune(version)
		runes[0] = unicode.ToUpper(runes[0])
		return string(runes)
	}
	return version
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

			caseStr := fmt.Sprintf("case *%s%s:\n\t\tv.ApiVersion = String(\"%s\")\n\t\tv.Kind = String(\"%s\")\n", gvk.Kind, capitalizedVersion, apiVersion, gvk.Kind)
			cases = append(cases, caseStr)

			testCaseStr := fmt.Sprintf(`{
				name:       "%s%s",
				input:      &%s%s{},
				wantApiVer: "%s",
				wantKind:   "%s",
			},`, gvk.Kind, capitalizedVersion, gvk.Kind, capitalizedVersion, apiVersion, gvk.Kind)
			testCases = append(testCases, testCaseStr)

			testAssertionStr := fmt.Sprintf(`case *%s%s:
				assert.Equal(t, tt.wantApiVer, *obj.ApiVersion)
				assert.Equal(t, tt.wantKind, *obj.Kind)
`, gvk.Kind, capitalizedVersion)
			testAssertions = append(testAssertions, testAssertionStr)
		}
	}

	// Sort the cases by kind
	sort.Strings(cases)
	sort.Strings(testCases)
	sort.Strings(testAssertions)

	switchFunction := "package k8s\n\nfunc SetApiVersionAndKind(i any) {\n\tswitch v := i.(type) {\n"
	for _, caseStr := range cases {
		switchFunction += "\t" + caseStr
	}
	switchFunction += "\t}\n}\n"

	testFunction := `package k8s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetApiVersionAndKind(t *testing.T) {
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
			SetApiVersionAndKind(tt.input)

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

func main() {
	version := flag.String("version", "", "Kubernetes version")
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
}
