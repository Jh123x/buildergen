package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestParseCommand_Fail(t *testing.T) {
	usageBuilder := strings.Builder{}
	cmd.GetUsage(func(s string, a ...any) (int, error) {
		usageBuilder.WriteString(fmt.Sprintf(s, a...))
		usageBuilder.WriteRune('\n')
		return 0, nil
	})

	tests := map[string]struct {
		src            string
		dest           string
		pkg            string
		name           string
		withValidation bool
		astMode        consts.Mode

		expectedOutput string
	}{
		"empty source": {
			src:            "",
			dest:           "test_destination.go",
			pkg:            "test_pkg",
			name:           "Test",
			withValidation: false,
			astMode:        consts.MODE_AST,
			expectedOutput: strings.Join(
				[]string{usageBuilder.String(), "Error parsing config file: " + consts.ErrSrcNotFound.Error(), "\n"},
				"",
			),
		},
		"invalid syntax": {
			src:            "../../go.mod",
			dest:           "test_destination.go",
			pkg:            "test_pkg",
			name:           "Test",
			withValidation: true,
			astMode:        consts.MODE_AST,
			expectedOutput: strings.Join(
				[]string{usageBuilder.String(), "Error parsing config file: " + consts.ErrNotGoFile.Error(), "\n"},
				"",
			),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			builder := strings.Builder{}
			tmp := func(s string, a ...any) (int, error) {
				builder.WriteString(fmt.Sprintf(s, a...))
				builder.WriteRune('\n')
				return 0, nil
			}

			ParseCommand(tc.src, tc.dest, tc.pkg, tc.name, tc.withValidation, string(tc.astMode), tmp)
			assert.Equal(t, tc.expectedOutput, builder.String())
		})
	}
}

func TestParseCommand_Success(t *testing.T) {
	tests := map[string]struct {
		src            string
		pkg            string
		name           string
		withValidation bool
		astMode        consts.Mode

		expectedFileRes string
	}{
		"simple struct": {
			src:             filepath.Join("data", "nest.go"),
			pkg:             "data",
			name:            "Test",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("data", "nest_expected_result.go"),
		},
		"struct with first letter cap": {
			src:             filepath.Join("data", "name_collision.go"),
			pkg:             "data",
			name:            "NameCollide",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("data", "name_expected_result.go"),
		},
		"struct keyword": {
			src:             filepath.Join("data", "keywords.go"),
			pkg:             "data",
			name:            "Struct",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("data", "keywords_expected.go"),
		},
		"benchmark struct": {
			src:             filepath.Join("..", "..", "examples", "benchmark", "benchmark.go"),
			pkg:             "benchmark",
			name:            "Data",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("..", "..", "examples", "benchmark", "benchmark_builder.go"),
		},
		"internal file": {
			src:             filepath.Join("..", "cmd", "ttypes.go"),
			pkg:             "cmd",
			name:            "Config",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder.go"),
		},
		"internal file test": {
			src:             filepath.Join("..", "cmd", "ttypes_test.go"),
			pkg:             "cmd",
			name:            "testCase",
			withValidation:  false,
			astMode:         consts.MODE_AST,
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder_test.go"),
		},
	}

	writeDir, err := os.MkdirTemp(consts.DEFAULT_TEMP_DIR, "test_parse_command")
	if !assert.Nil(t, err) {
		return
	}

	t.Cleanup(func() { assert.Nil(t, os.RemoveAll(writeDir)) })

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			dest := filepath.Join(writeDir, fmt.Sprintf("%s.go", name))
			ParseCommand(tc.src, dest, tc.pkg, tc.name, tc.withValidation, string(tc.astMode), fmt.Printf)

			destRes, err := os.ReadFile(tc.expectedFileRes)
			assert.Nil(t, err)

			expectedRes := consts.EMPTY_STR

			if len(tc.expectedFileRes) == 0 {
				assert.Equal(t, consts.EMPTY_STR, string(destRes))
				return
			}

			rawRes, err := os.ReadFile(tc.expectedFileRes)
			assert.Nil(t, err)
			expectedRes = string(rawRes)
			assert.Equal(t, expectedRes, string(destRes))
		})
	}
}
