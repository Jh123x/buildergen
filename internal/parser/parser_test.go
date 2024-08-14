package parser

import (
	"os"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
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
				Source:      currDir + "/data/nest.go",
				Destination: currDir + "/data/nest_result.go",
				Package:     "nested",
				Name:        "Test",
			},
			expectedFileRes: currDir + "/data/nest_expected_result.go",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := ParseBuilderFile(tc.config)
			expectedRes := ""

			if len(tc.expectedFileRes) > 0 {
				rawRes, err := os.ReadFile(tc.expectedFileRes)
				assert.Nil(t, err)
				expectedRes = string(rawRes)
			}

			assert.Equal(t, expectedRes, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
