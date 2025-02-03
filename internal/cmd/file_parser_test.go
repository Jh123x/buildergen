package cmd

import (
	"io/fs"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestParseConfigFile(t *testing.T) {
	tests := map[string]struct {
		configPath      string
		expectedConfigs []*Config
		expectedErr     error
	}{
		"empty path": {
			configPath:  "",
			expectedErr: consts.ErrInvalidConfigFile,
		},
		"file not found": {
			configPath: "not found",
			expectedErr: &fs.PathError{
				Op:   "open",
				Path: "not found",
				Err:  syscall.Errno(2),
			},
		},
		"invalid yaml": {
			configPath: filepath.Join("..", "..", "go.mod"),
			expectedErr: &yaml.TypeError{
				Errors: []string{"line 1: cannot unmarshal !!str `module ...` into cmd.BuilderGenConfig"},
			},
		},
		"error filling up default": {
			configPath:  filepath.Join("data", "invalid.yaml"),
			expectedErr: consts.ErrNotGoFile,
		},
		"valid config": {
			configPath: filepath.Join("..", "..", ".builder.yaml"),
			expectedConfigs: []*Config{
				{
					Source:      "./internal/cmd/ttypes.go",
					Destination: "./internal/cmd/ttypes_builder.go",
					Name:        "Config",
					ParserMode:  consts.MODE_AST,
				},
				{
					Source:      "./internal/cmd/ttypes_test.go",
					Destination: "./internal/cmd/ttypes_builder_test.go",
					Name:        "testCase",
					ParserMode:  consts.MODE_AST,
				},
				{
					Source:      "./examples/benchmark/benchmark.go",
					Destination: "./examples/benchmark/benchmark_builder.go",
					Name:        "Data",
					ParserMode:  consts.MODE_AST,
				},
				{
					Source:      "./examples/person.go",
					Destination: "./examples/person_builder.go",
					Name:        "Person",
					ParserMode:  consts.MODE_FAST,
				},
				{
					Source:      "./examples/person.go",
					Destination: "./examples/person_builder.go",
					Name:        "UnRelated",
					ParserMode:  consts.MODE_FAST,
				},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := ParseConfigFile(tc.configPath)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedConfigs, res)
		})
	}
}
