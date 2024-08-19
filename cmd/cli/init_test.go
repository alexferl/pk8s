package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCmd(t *testing.T) {
	t.Run("init in empty directory", func(t *testing.T) {
		tempDir := t.TempDir()

		exec := NewMockExecutor(t)
		exec.EXPECT().Init(tempDir).Return([]byte{}, nil)
		exec.EXPECT().Tidy(tempDir).Return([]byte{}, nil)

		cmd := initCmd(exec)
		var out bytes.Buffer
		cmd.SetOut(&out)
		cmd.SetArgs([]string{tempDir})
		err := cmd.Execute()
		assert.NoError(t, err)

		assert.Equal(t, initMsg, out.String())
	})

	t.Run("init in non-empty directory", func(t *testing.T) {
		tempDir := t.TempDir()

		file, err := os.Create(filepath.Join(tempDir, "file.txt"))
		assert.NoError(t, err)
		file.Close()

		cmd := initCmd(NewMockExecutor(t))
		var out bytes.Buffer
		cmd.SetOut(&out)
		cmd.SetArgs([]string{tempDir})
		err = cmd.Execute()
		assert.Error(t, err)

		expectedErrMsg := "directory '" + tempDir + "' is not empty\n"
		assert.Contains(t, out.String(), expectedErrMsg)
	})

	t.Run("init in non-existent directory", func(t *testing.T) {
		tempDir := filepath.Join(t.TempDir(), "nonexistent")

		exec := NewMockExecutor(t)
		exec.EXPECT().Init(tempDir).Return([]byte{}, nil)
		exec.EXPECT().Tidy(tempDir).Return([]byte{}, nil)

		cmd := initCmd(exec)
		var out bytes.Buffer
		cmd.SetOut(&out)
		cmd.SetArgs([]string{tempDir})
		err := cmd.Execute()
		assert.NoError(t, err)

		assert.Equal(t, initMsg, out.String())
	})
}
