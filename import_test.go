package pk8s

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestImporterRead(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		setupMock   func()
		cleanupMock func()
		wantData    string
		wantErr     bool
	}{
		{
			name: "Custom name",
			path: "custom:=testfile.yaml",
			setupMock: func() {
				os.WriteFile("testfile.yaml", []byte("test data"), 0o644)
			},
			cleanupMock: func() {
				os.Remove("testfile.yaml")
			},
			wantData: "test data",
			wantErr:  false,
		},
		{
			name: "Stdin input",
			path: "-",
			setupMock: func() {
				oldStdin := os.Stdin
				r, w, _ := os.Pipe()
				os.Stdin = r
				go func() {
					w.Write([]byte("stdin data"))
					w.Close()
				}()
				t.Cleanup(func() { os.Stdin = oldStdin })
			},
			wantData: "stdin data",
			wantErr:  false,
		},
		{
			name: "HTTP link",
			path: "http://example.com/data",
			setupMock: func() {
				oldClient := http.DefaultClient
				http.DefaultClient = &http.Client{
					Transport: &mockTransport{
						response: &http.Response{
							StatusCode: 200,
							Body:       io.NopCloser(strings.NewReader("http data")),
						},
					},
				}
				t.Cleanup(func() { http.DefaultClient = oldClient })
			},
			wantData: "http data",
			wantErr:  false,
		},
		{
			name: "File path",
			path: "testfile.yaml",
			setupMock: func() {
				os.WriteFile("testfile.yaml", []byte("file data"), 0o644)
			},
			cleanupMock: func() {
				os.Remove("testfile.yaml")
			},
			wantData: "file data",
			wantErr:  false,
		},
		{
			name:    "Non-existent file",
			path:    "nonexistent.yaml",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setupMock != nil {
				tt.setupMock()
			}
			if tt.cleanupMock != nil {
				defer tt.cleanupMock()
			}

			i := &importer{}
			gotData, err := i.Read(tt.path)

			if (err != nil) != tt.wantErr {
				t.Errorf("importer.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(gotData) != tt.wantData {
				t.Errorf("importer.Read() = %v, want %v", string(gotData), tt.wantData)
			}
		})
	}
}

// mockTransport is a mock for http.RoundTripper
type mockTransport struct {
	response *http.Response
}

func (m *mockTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return m.response, nil
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

	importer := NewImporter(&ImporterConfig{Overwrite: true})

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
