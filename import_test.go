package pk8s

import (
	"bytes"
	"os"
	"testing"
)

func TestImporterRead(t *testing.T) {
	content := []byte("test content")
	tmpFile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatal(err)
	}

	importer := NewImporter(&ImporterConfig{})

	data, err := importer.Read(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if !bytes.Equal(data, content) {
		t.Errorf("Read incorrect content: got %q, want %q", data, content)
	}

	// TODO:
	// Test reading from stdin
	// ...

	// Test reading from HTTP
	// ...
}

func TestImporterImport(t *testing.T) {
	crdYAML := `
apiVersion: apiextensions.k8s.io/v1
kind: CustomResourceDefinition
metadata:
  name: testcrds.example.com
spec:
  group: example.com
  names:
    kind: TestCRD
    listKind: TestCRDList
    plural: testcrds
    singular: testcrd
  scope: Namespaced
  versions:
  - name: v1
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              field1:
                type: string
              field2:
                type: integer
`

	importer := &importer{
		config: &ImporterConfig{Overwrite: true},
	}

	err := importer.Import([]byte(crdYAML))
	if err != nil {
		t.Fatalf("Failed to import CRD: %v", err)
	}

	_, err = os.Stat("imports/example_com/test_crd.go")
	if os.IsNotExist(err) {
		t.Errorf("Expected file was not created")
	}

	// TODO: verifying the content of the created file
}

func TestHasCustomName(t *testing.T) {
	tests := []struct {
		input     string
		wantName  string
		wantValue string
		wantOk    bool
	}{
		{"custom:=value", "custom", "value", true},
		{"normal_input", "", "", false},
		{"custom:=", "custom", "", true},
	}

	for _, tt := range tests {
		gotName, gotValue, gotOk := hasCustomName(tt.input)
		if gotName != tt.wantName || gotValue != tt.wantValue || gotOk != tt.wantOk {
			t.Errorf("hasCustomName(%q) = (%q, %q, %v), want (%q, %q, %v)",
				tt.input, gotName, gotValue, gotOk, tt.wantName, tt.wantValue, tt.wantOk)
		}
	}
}

func TestIsHTTPLink(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"http://example.com", true},
		{"https://example.com", true},
		{"ftp://example.com", false},
		{"not_a_link", false},
	}

	for _, tt := range tests {
		got := isHTTPLink(tt.input)
		if got != tt.want {
			t.Errorf("isHTTPLink(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestGetGoType(t *testing.T) {
	tests := []struct {
		name string
		prop Property
		want string
	}{
		{"String", Property{Type: "string"}, "string"},
		{"Integer", Property{Type: "integer"}, "int"},
		{"Boolean", Property{Type: "boolean"}, "bool"},
		{"Array", Property{Type: "array", Items: &Property{Type: "string"}}, "[]string"},
		{"Object", Property{Type: "object", Properties: map[string]Property{}}, "map[string]interface{}"},
		{"CustomObject", Property{Type: "object", Properties: map[string]Property{"field": {}}}, "CustomObject"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getGoType(tt.name, tt.prop)
			if got != tt.want {
				t.Errorf("getGoType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateStructs(t *testing.T) {
	prop := Property{
		Type: "object",
		Properties: map[string]Property{
			"field1": {Type: "string"},
			//"field2": {Type: "integer"},
		},
	}

	var file File
	generateStructs(&file, "TestStruct", prop)

	expected := `type TestStruct struct {
	Field1 *string ` + "`json:\"field1,omitempty\"`" + `
}

`

	if file.String() != expected {
		t.Errorf("generateStructs() generated incorrect struct:\n%s\nExpected:\n%s", file.String(), expected)
	}
}
