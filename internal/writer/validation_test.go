package writer

import (
	"errors"
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
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
						{Name: "test7", Path: `"test.com/test/test"`},
					},
				},
				{
					Imports: []*generation.Import{
						{Name: "test7", Path: `"test.com/test/test"`},
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
				{Name: "test7", Path: `"test.com/test/test"`},
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

func Test_mergePackage(t *testing.T) {
	tests := map[string]struct {
		structs     []*generation.StructGenHelper
		expectedRes string
		expectedErr error
	}{
		"empty": {
			structs:     nil,
			expectedRes: consts.EMPTY_STR,
			expectedErr: nil,
		},
		"no package found": {
			structs: []*generation.StructGenHelper{
				{Name: "test"},
				{Name: "test2"},
			},
			expectedRes: consts.EMPTY_STR,
			expectedErr: errors.New("no packages found within structs: test, test2"),
		},
		"only 1 struct": {
			structs: []*generation.StructGenHelper{
				{
					Name:    "test",
					Package: "test_package",
				},
			},
			expectedRes: "test_package",
		},
		"multiple non-conflict package": {
			structs: []*generation.StructGenHelper{
				{Name: "test", Package: "test_package"},
				{Name: "test2", Package: "test_package"},
				{Name: "test3", Package: "test_package"},
			},
			expectedRes: "test_package",
		},
		"multiple conflict package": {
			structs: []*generation.StructGenHelper{
				{Name: "test", Package: "test_package"},
				{Name: "test2", Package: "test_package2"},
				{Name: "test3", Package: "test_package"},
			},
			expectedRes: consts.EMPTY_STR,
			expectedErr: errors.New("multiple packages found within the same file: test_package, test_package2"),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, err := mergePackages(tc.structs)
			assert.Equal(t, tc.expectedErr, err)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
