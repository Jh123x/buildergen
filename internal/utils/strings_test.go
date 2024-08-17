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

func TestAny(t *testing.T) {
	tests := map[string]struct {
		pred        func(string) bool
		data        []string
		expectedRes bool
	}{
		"match all predicate should return true": {
			pred:        func(s string) bool { return true },
			data:        []string{"data1", "data2", "data3"},
			expectedRes: true,
		},
		"match 1 predicate should return true": {
			pred:        func(s string) bool { return s == "test" },
			data:        []string{"t", "te", "tes", "test"},
			expectedRes: true,
		},
		"match non should return false": {
			pred:        func(s string) bool { return false },
			data:        []string{"data1", "data2", "data3"},
			expectedRes: false,
		},
		"empty arr should return false": {
			pred:        func(s string) bool { return true },
			data:        []string{},
			expectedRes: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, Any(tc.pred, tc.data...))
		})
	}
}
