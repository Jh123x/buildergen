package writer

import (
	"errors"
	"testing"

	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/stretchr/testify/assert"
)

func Test_mergeImports(t *testing.T) {
	tests := map[string]struct {
		structs         []*generation.StructGenHelper
		expectedImports []*generation.Import
		expectedErr     error
	}{
		"empty structs": {
			structs:         []*generation.StructGenHelper{},
			expectedImports: []*generation.Import{},
			expectedErr:     nil,
		},
		"only 1 struct": {
			structs: []*generation.StructGenHelper{
				{
					Imports: []*generation.Import{
						{Path: `"test.com/test/test"`},
					},
				},
			},
			expectedImports: []*generation.Import{
				{Path: `"test.com/test/test"`},
			},
		},
		"multiple non-conflict should combine": {
			structs: []*generation.StructGenHelper{
				{
					Imports: []*generation.Import{
						{Name: "test", Path: `"test.com/test/test"`},
					},
				},
				{
					Imports: []*generation.Import{
						{Path: `"test.com/test/test"`},
						{Path: `"test.com/test/test2"`},
					},
				},
				{
					Imports: []*generation.Import{
						{Path: `"test.com/test/test3"`},
						{Path: `"test.com/test/test5"`},
					},
				},
			},
			expectedImports: []*generation.Import{
				{Path: `"test.com/test/test"`},
				{Path: `"test.com/test/test2"`},
				{Path: `"test.com/test/test3"`},
				{Path: `"test.com/test/test5"`},
			},
		},
		"multiple conflict should error": {
			structs: []*generation.StructGenHelper{
				{
					Name: "test",
					Imports: []*generation.Import{
						{Name: "test_conf", Path: `"test.com/test/test"`},
					},
				},
				{
					Name: "test1",
					Imports: []*generation.Import{
						{Path: `"test.com/test/test"`},
						{Path: `"test.com/test/test2"`},
					},
				},
				{
					Name: "test2",
					Imports: []*generation.Import{
						{Path: `"test.com/test/test3"`},
						{Path: `"test.com/test/test5"`},
					},
				},
			},
			expectedImports: nil,
			expectedErr:     errors.New("import names for package='\"test.com/test/test\"' are inconsistent between test (test_conf) and test1 (test)"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := mergeImports(tc.structs)
			assert.Equal(t, tc.expectedErr, err)
			assert.ElementsMatch(t, tc.expectedImports, res)
		})
	}
}
