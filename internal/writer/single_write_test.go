package writer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteToSingleFile(t *testing.T) {
	tests := map[string]struct {
		dest string
		code string
	}{
		"write empty file": {
			dest: "test_file.txt",
			code: "",
		},
		"write non-empty file": {
			dest: "test_file.txt",
			code: "some stuff",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tmpDir := os.TempDir()
			writeDir := filepath.Join(tmpDir, "test_dir")

			assert.Nil(t, os.Mkdir(filepath.Join(tmpDir, "test_dir"), 0644))
			defer func() { assert.Nil(t, os.RemoveAll(writeDir)) }()
			filePath := filepath.Join(writeDir, tc.dest)
			assert.Nil(t, WriteToSingleFile(filePath, tc.code))

			data, err := os.ReadFile(filePath)
			assert.Nil(t, err)
			assert.Equal(t, tc.code, string(data))
		})
	}
}
