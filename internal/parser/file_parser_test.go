package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	usageBuilder := strings.Builder{}
	cmd.GetUsage(func(s string, a ...any) (int, error) {
		usageBuilder.WriteString(fmt.Sprintf(s, a...))
		usageBuilder.WriteRune('\n')
		return 0, nil
	})

	tests := map[string]struct {
		configs         []*cmd.Config
		expectedFileRes string
	}{
		"simple struct": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("data", "nest.go"),
					Package:        "data",
					Name:           "Test",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("data", "nest_expected_result.go"),
		},
		"struct with first letter cap": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("data", "name_collision.go"),
					Package:        "data",
					Name:           "NameCollide",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("data", "name_expected_result.go"),
		},
		"struct keyword": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("data", "keywords.go"),
					Package:        "data",
					Name:           "Struct",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("data", "keywords_expected.go"),
		},
		"benchmark struct": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("..", "..", "examples", "benchmark", "benchmark.go"),
					Package:        "benchmark",
					Name:           "Data",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("..", "..", "examples", "benchmark", "benchmark_builder.go"),
		},
		"internal file": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("..", "cmd", "ttypes.go"),
					Package:        "cmd",
					Name:           "Config",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder.go"),
		},
		"internal file test": {
			configs: []*cmd.Config{
				{
					Source:         filepath.Join("..", "cmd", "ttypes_test.go"),
					Package:        "cmd",
					Name:           "testCase",
					WithValidation: false,
					ParserMode:     consts.MODE_AST,
				},
			},
			expectedFileRes: filepath.Join("..", "cmd", "ttypes_builder_test.go"),
		},
	}

	writeDir, err := os.MkdirTemp(consts.DEFAULT_TEMP_DIR, "test_parse_and_write_builder_file")
	if !assert.Nil(t, err) {
		return
	}

	t.Cleanup(func() { assert.Nil(t, os.RemoveAll(writeDir)) })

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			builder := strings.Builder{}
			defer builder.Reset()
			tmp := func(s string, a ...any) (int, error) {
				builder.WriteString(fmt.Sprintf(s, a...))
				builder.WriteRune('\n')
				return 0, nil
			}

			dest := filepath.Join(writeDir, fmt.Sprintf("%s.go", name))
			tc.configs = utils.Map(
				tc.configs,
				func(c *cmd.Config) *cmd.Config {
					c.Destination = dest
					return c
				},
			)

			ParseAndWriteBuilderFile(tc.configs, tmp)
			if len(tc.expectedFileRes) == 0 {
				return
			}

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
