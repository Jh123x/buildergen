package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	tests := map[string]struct {
		lst         []int
		predicate   func(int) bool
		expectedRes []int
	}{
		"empty": {
			lst:         make([]int, 0),
			predicate:   func(i int) bool { return true },
			expectedRes: make([]int, 0),
		},
		"filter out no values": {
			lst:         []int{1, 2, 3, 4, 5, 6, 7, 8},
			predicate:   func(i int) bool { return true },
			expectedRes: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		"filter out values": {
			lst:         []int{1, 2, 3, 4, 5, 6, 7, 8},
			predicate:   func(i int) bool { return i%2 == 0 },
			expectedRes: []int{2, 4, 6, 8},
		},
		"filter out all values": {
			lst:         []int{1, 2, 3, 4, 5, 6, 7, 8},
			predicate:   func(i int) bool { return i%2 == 2 },
			expectedRes: []int{},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, Filter(tc.lst, tc.predicate))
		})
	}
}

func TestMap(t *testing.T) {
	tests := map[string]struct {
		lst         []string
		mapper      func(string) string
		expectedRes []string
	}{
		"empty": {
			lst:         nil,
			mapper:      func(s string) string { return s + s },
			expectedRes: []string{},
		},
		"no mapper": {
			lst:         []string{"1", "2", "3"},
			mapper:      func(s string) string { return s },
			expectedRes: []string{"1", "2", "3"},
		},
		"not empty": {
			lst:         []string{"1", "2", "3"},
			mapper:      func(s string) string { return s + s },
			expectedRes: []string{"11", "22", "33"},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, Map(tc.lst, tc.mapper))
		})
	}
}

func TestFilterNil(t *testing.T) {
	no1 := 0
	no2 := 1

	tests := map[string]struct {
		lst         []*int
		expectedRes []*int
	}{
		"empty": {
			lst:         nil,
			expectedRes: []*int{},
		},
		"has nil": {
			lst:         []*int{&no1, &no2, nil, &no1, nil, &no2},
			expectedRes: []*int{&no1, &no2, &no1, &no2},
		},
		"has no nil": {
			lst:         []*int{&no1, &no2, &no1, &no2},
			expectedRes: []*int{&no1, &no2, &no1, &no2},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expectedRes, FilterNil(tc.lst))
		})
	}
}
