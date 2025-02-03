package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSet(t *testing.T) {
	defer func() { assert.Nil(t, recover()) }()
	s := NewSet[string]()
	assert.NotNil(t, s)
}

func TestSet(t *testing.T) {
	tests := map[string]struct {
		opts        func(s *Set[string])
		testKey     string
		expectedHas bool
	}{
		"add empty string should retrieve empty string": {
			opts: func(s *Set[string]) {
				s.Add("")
			},
			testKey:     "",
			expectedHas: true,
		},
		"add other string empty string should be false": {
			testKey:     "",
			expectedHas: false,
		},
		"add key twice should still be true": {
			testKey: "test",
			opts: func(s *Set[string]) {
				s.Add("test")
				s.Add("test")
			},
			expectedHas: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := NewSet[string]()
			if tc.opts != nil {
				tc.opts(&s)
			}
			assert.Equal(t, tc.expectedHas, s.Has(tc.testKey))
		})
	}
}

func TestSet_ToList(t *testing.T) {
	tests := map[string]struct {
		s           Set[string]
		expectedLst []string
	}{
		"empty set": {
			s:           NewSet[string](),
			expectedLst: []string{},
		},
		"filled set": {
			s:           NewSet[string]("a", "b", "c", "c", "d"),
			expectedLst: []string{"a", "b", "c", "d"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, len(tc.expectedLst), tc.s.Len())
			assert.ElementsMatch(t, tc.expectedLst, tc.s.ToList())
		})
	}
}
