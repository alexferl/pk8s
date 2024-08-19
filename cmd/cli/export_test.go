package cli

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExportCmd(t *testing.T) {
	t.Run("failed command execution", func(t *testing.T) {
		cmd := exportCmd()
		var out bytes.Buffer
		cmd.SetOut(&out)
		cmd.SetArgs([]string{"."})
		err := cmd.Execute()
		assert.Error(t, err)

		assert.Contains(t, out.String(), "command execution failed")
	})
}
