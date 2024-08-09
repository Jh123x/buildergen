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

func TestContains(t *testing.T) {
	tests := map[string]struct {
		arr         []string
		item        string
		expectedRes bool
	}{
		"item found should return true": {
			arr:         []string{"test", "test1", "test2"},
			item:        "test1",
			expectedRes: true,
		},
		"empty arr should return false": {
			arr:         []string{},
			item:        "test",
			expectedRes: false,
		},
		"item not found should return false": {
			arr:         []string{"test1", "test2", "test3"},
			item:        "test4",
			expectedRes: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, Contains(tc.arr, tc.item))
		})
	}
}
