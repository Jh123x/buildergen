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
