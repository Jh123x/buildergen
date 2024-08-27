package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNilOrEmpty(t *testing.T) {
	testEmptyStr := ""
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
