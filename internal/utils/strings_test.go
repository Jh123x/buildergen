package utils

import (
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestIsNilOrEmpty(t *testing.T) {
	testEmptyStr := consts.EMPTY_STR
	testStr := "t"

	tests := map[string]struct {
		val         *string
		expectedRes bool
	}{
		"nil should return true": {
			val:         nil,
			expectedRes: true,
		},
		"empty string should return true": {
			val:         &testEmptyStr,
			expectedRes: true,
		},
		"non empty string should return true": {
			val:         &testStr,
			expectedRes: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, IsNilOrEmpty(tc.val))
		})
	}
}

func TestLowerFirstLetter(t *testing.T) {
	tests := map[string]struct {
		val         string
		expectedVal string
	}{
		"empty": {
			val:         "",
			expectedVal: "",
		},
		"caps first letter": {
			val:         "ABC",
			expectedVal: "aBC",
		},
		"small first letter": {
			val:         "abC",
			expectedVal: "abC",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedVal, LowerFirstLetter(tc.val))
		})
	}
}
