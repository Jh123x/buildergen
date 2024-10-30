package parser

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestParseBuilderFile(t *testing.T) {
	currDir, err := os.Getwd()
	assert.Nil(t, err)

	tests := map[string]struct {
		config          *cmd.Config
		expectedFileRes string
		expectedErr     error
	}{
		"simple struct": {
			config: &cmd.Config{
				Source:  currDir + "/data/nest.go",
				Package: "data",
				Name:    "Test",
			},
			expectedFileRes: currDir + "/data/nest_expected_result.go",
		},
		"struct with first letter cap": {
			config: &cmd.Config{
				Source:  currDir + "/data/name_collision.go",
				Package: "data",
				Name:    "NameCollide",
			},
			expectedFileRes: currDir + "/data/name_expected_result.go",
		},
		"struct keyword": {
			config: &cmd.Config{
				Source:  currDir + "/data/keywords.go",
				Package: "data",
				Name:    "Struct",
			},
			expectedFileRes: currDir + "/data/keywords_expected.go",
		},
		"benchmark struct": {
			config: &cmd.Config{
				Source:  path.Join(currDir, "..", "..", "examples", "benchmark", "benchmark.go"),
				Package: "benchmark",
				Name:    "Data",
			},
			expectedFileRes: path.Join(currDir, "..", "..", "examples", "benchmark", "benchmark_builder.go"),
		},
		"internal file": {
			config: &cmd.Config{
				Source:  path.Join(currDir, "..", "cmd", "ttypes.go"),
				Package: "cmd",
				Name:    "Config",
			},
			expectedFileRes: path.Join(currDir, "..", "cmd", "ttypes_builder.go"),
		},
		"internal file test": {
			config: &cmd.Config{
				Source:  path.Join(currDir, "..", "cmd", "ttypes_test.go"),
				Package: "cmd",
				Name:    "testCase",
			},
			expectedFileRes: path.Join(currDir, "..", "cmd", "ttypes_builder_test.go"),
		},
	}

	for _, mode := range consts.ALL_MODES {
		for name, tc := range tests {
			t.Run(fmt.Sprintf("%s_%s", name, mode), func(t *testing.T) {
				if tc.config != nil {
					tc.config.ParserMode = mode
				}

				res, err := ParseBuilderFile(tc.config)
				expectedRes := consts.EMPTY_STR

				if len(tc.expectedFileRes) > 0 {
					rawRes, err := os.ReadFile(tc.expectedFileRes)
					assert.Nil(t, err)
					expectedRes = string(rawRes)
				}

				assert.Equal(t, tc.expectedErr, err)
				assert.Equal(t, expectedRes, res)
			})
		}
	}
}
