package cmd

import (
	"bytes"
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	result := example()
	if !strings.Contains(result, "pk8s import cert-manager.crds.yaml") {
		t.Error("example() output doesn't contain expected content")
	}
	// Add more checks for other expected content
}

func TestImportCmd(t *testing.T) {
	// Test case 1: No arguments
	//cmd := importCmd()
	//buf := new(bytes.Buffer)
	//cmd.SetOut(buf)
	//cmd.SetArgs([]string{})
	//err := cmd.Execute()
	//if err != nil {
	//	t.Errorf("Execute() with no args returned unexpected error: %v", err)
	//}

	// Test case 2: With argument
	cmd := importCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"test.yaml"})
	err := cmd.Execute()
	if err == nil {
		t.Error("Execute() with non-existent file should return an error")
	}

	// Test case 3: With overwrite flag
	cmd = importCmd()
	buf = new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--overwrite", "test.yaml"})
	err = cmd.Execute()
	if err == nil {
		t.Error("Execute() with non-existent file should return an error, even with overwrite flag")
	}
}
