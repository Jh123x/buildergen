package writer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestWriteToSingleFile(t *testing.T) {
	tests := map[string]struct {
		dest string
		code string
	}{
		"write empty file": {
			dest: "test_file_empty.txt",
			code: "",
		},
		"write non-empty file": {
			dest: "test_file_non_empty.txt",
			code: "some stuff",
		},
	}

	writeDir, err := os.MkdirTemp(consts.DEFAULT_TEMP_DIR, "test_write_to_single_file")
	if !assert.Nil(t, err) {
		return
	}

	t.Cleanup(func() { assert.Nil(t, os.RemoveAll(writeDir)) })

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			filePath := filepath.Join(writeDir, tc.dest)
			if !assert.Nil(t, WriteToSingleFile(filePath, tc.code), "Error writing to file") {
				return
			}

			data, err := os.ReadFile(filePath)
			assert.Nil(t, err, "Error reading from file")
			assert.Equal(t, tc.code, string(data))
		})
	}
}
