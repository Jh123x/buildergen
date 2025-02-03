package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestParseBuilderFile(t *testing.T) {
	tests := map[string]struct {
		config          *cmd.Config
		expectedFileRes string
		expectedErr     error
	}{
		"simple struct": {
			config: &cmd.Config{
				Source:  filepath.Join("data", "nest.go"),
				Package: "data",
				Name:    "Test",
			},
			expectedFileRes: filepath.Join("data", "nest_expected_result.go"),
		},
		"struct with first letter cap": {
			config: &cmd.Config{
				Source:  filepath.Join("data", "name_collision.go"),
				Package: "data",
				Name:    "NameCollide",
			},
			expectedFileRes: filepath.Join("data", "name_expected_result.go"),
		},
		"struct keyword": {
			config: &cmd.Config{
				Source:  filepath.Join("data", "keywords.go"),
				Package: "data",
				Name:    "Struct",
			},
			expectedFileRes: filepath.Join("data", "keywords_expected.go"),
		},
		"benchmark struct": {
			config: &cmd.Config{
				Source:  filepath.Join("..", "..", "examples", "benchmark", "benchmark.go"),
				Package: "benchmark",
				Name:    "Data",
			},
			expectedFileRes: filepath.Join("..", "..", "examples", "benchmark", "benchmark_builder.go"),
		},
		"internal file": {
			config: &cmd.Config{
				Source:  filepath.Join("..", "cmd", "ttypes.go"),
				Package: "cmd",
				Name:    "Config",
			},
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder.go"),
		},
		"internal file test": {
			config: &cmd.Config{
				Source:  filepath.Join("..", "cmd", "ttypes_test.go"),
				Package: "cmd",
				Name:    "testCase",
			},
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder_test.go"),
		},
	}

	currDir, err := os.Getwd()
	assert.Nil(t, err)

	for name, tc := range tests {
		if tc.config != nil {
			tc.config.Source = filepath.Join(currDir, tc.config.Source)
			tc.expectedFileRes = filepath.Join(currDir, tc.expectedFileRes)
		}

		t.Log(tc.config.Source)
		t.Log(tc.expectedFileRes)

		for _, mode := range consts.ALL_MODES {
			t.Run(fmt.Sprintf("%s_%s", name, mode), func(t *testing.T) {
				tc.config.ParserMode = mode
				res, err := ParseBuilderFile(tc.config)
				expectedRes := consts.EMPTY_STR

				if len(tc.expectedFileRes) > 0 {
					rawRes, err := os.ReadFile(tc.expectedFileRes)
					assert.Nil(t, err)
					expectedRes = string(rawRes)
				}

				assert.Equal(t, tc.expectedErr, err)
				assert.Equal(t, expectedRes, res.ToSource())
			})
		}
	}
}

func Test_getParserMode(t *testing.T) {
	tests := map[string]struct {
		parserMode  consts.Mode
		expectedRes parserFn
	}{
		"ast mode": {
			parserMode:  consts.MODE_AST,
			expectedRes: parseDataByAST,
		},
		"custom parser": {
			parserMode:  consts.MODE_FAST,
			expectedRes: parseDataByCustomParser,
		},
		"not found": {
			parserMode:  "not found mode",
			expectedRes: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			sf1 := reflect.ValueOf(tc.expectedRes)
			sf2 := reflect.ValueOf(getParserMode(tc.parserMode))

			assert.Equal(t, sf1.Pointer(), sf2.Pointer())
		})
	}
}
