package writer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/stretchr/testify/assert"
)

func TestMultiFileWrite(t *testing.T) {
	tests := map[string]struct {
		path            string
		structs         []*generation.StructGenHelper
		expectedResPath string
		expectedErr     error
	}{
		"1 simple struct": {
			path: "simple_struct.go",
			structs: []*generation.StructGenHelper{
				nil,
				{
					Name:    "Test",
					Package: "data",
					Fields: []*generation.Field{
						{Name: "Val", Type: "string"},
						{Name: "ImportedType", Type: "*os.FileMode"},
					},
					Imports: []*generation.Import{
						{Path: `"os"`},
					},
				},
			},
			expectedResPath: filepath.Join("..", "parser", "data", "nest_expected_result.go"),
		},
		"2 simple struct": {
			path: "simple_struct.go",
			structs: []*generation.StructGenHelper{
				nil,
				{
					Name:    "Test",
					Package: "data",
					Fields: []*generation.Field{
						{Name: "Val", Type: "string"},
						{Name: "ImportedType", Type: "*os.FileMode"},
					},
					Imports: []*generation.Import{
						{Path: `"os"`},
					},
				},
				{
					Name:    "OtherStruct",
					Package: "data",
					Fields: []*generation.Field{
						{Name: "OtherImports", Type: "context.Context"},
					},
					Imports: []*generation.Import{
						{Path: `"os"`},
						{Path: `"context"`},
					},
				},
			},
			expectedResPath: filepath.Join("data", "simple_multi_result.go"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			tmpDir := os.TempDir()
			writeDir := filepath.Join(tmpDir, "test_dir")

			assert.Nil(t, os.Mkdir(filepath.Join(tmpDir, "test_dir"), 0644))
			defer func() { assert.Nil(t, os.RemoveAll(writeDir)) }()

			filePath := filepath.Join(writeDir, tc.path)
			assert.Equal(t, tc.expectedErr, MultiFileWrite(filePath, tc.structs...))

			if tc.expectedErr != nil {
				return
			}

			expectedData, err := os.ReadFile(tc.expectedResPath)
			assert.Nil(t, err)

			data, err := os.ReadFile(filePath)
			assert.Nil(t, err)

			assert.Equal(t, string(expectedData), string(data))
		})
	}
}
