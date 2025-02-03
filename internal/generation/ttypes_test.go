package generation

import (
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestImport_ToImport(t *testing.T) {
	tests := map[string]struct {
		iVal           *Import
		expectedResult string
	}{
		"import with no name": {
			iVal: &Import{
				Path: `"github.com/test/test"`,
			},
			expectedResult: `"github.com/test/test"`,
		},
		"import with name": {
			iVal: &Import{
				Name: "other",
				Path: `"github.com/test/test"`,
			},
			expectedResult: `other "github.com/test/test"`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.iVal.ToImport())
		})
	}
}

func TestImport_GetName(t *testing.T) {
	tests := map[string]struct {
		iVal           *Import
		expectedResult string
	}{
		"import with no name": {
			iVal: &Import{
				Path: `"github.com/test/test"`,
			},
			expectedResult: `test`,
		},
		"import with name": {
			iVal: &Import{
				Name: "other",
				Path: `"github.com/test/test"`,
			},
			expectedResult: `other`,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.iVal.GetName())
		})
	}
}

func TestField_GetUsedPackageName(t *testing.T) {
	tests := map[string]struct {
		field          Field
		expectedResult string
	}{
		"primitive type": {
			field: Field{
				Name: "test",
				Type: "int",
				Tags: "`json:\"test\"`",
			},
			expectedResult: consts.EMPTY_STR,
		},
		"non primitive type": {
			field: Field{
				Name: "test",
				Type: "*strings.Builder",
				Tags: "`json:\"test\"`",
			},
			expectedResult: "strings",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedResult, tc.field.GetUsedPackageName())
		})
	}
}
