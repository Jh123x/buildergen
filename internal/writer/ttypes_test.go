package writer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/stretchr/testify/assert"
)

func Test_writeHelper(t *testing.T) {
	tests := map[string]struct {
		writeHelper     *writeHelper
		expectedResPath string
	}{
		"single struct": {
			writeHelper: &writeHelper{
				pkg: "data",
				imports: []*generation.Import{
					{Path: `"strings"`},
					{Path: `"github.com/Jh123x/buildergen/internal/generation"`},
				},
				structs: []*generation.StructGenHelper{
					{
						Name:    "test",
						Package: "should be ignored",
						Fields: []*generation.Field{
							{Name: "test", Type: "generation.Field", Tags: ""},
							{Name: "test2", Type: "strings.Builder", Tags: ""},
						},
						Imports: []*generation.Import{
							{Name: "", Path: `"github.com/test/test"`},
						},
					},
				},
			},
			expectedResPath: filepath.Join(".", "data", "single_struct.txt"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := os.ReadFile(tc.expectedResPath)
			assert.Nil(t, err)
			assert.Equal(t, string(res), tc.writeHelper.ToSource())
		})
	}
}
