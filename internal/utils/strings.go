package utils

// IsNilOrEmpty returns true if the value is nil or equal to the empty value.
func IsNilOrEmpty[inputType comparable](val *inputType) bool {
	var empty inputType
	if val == nil || *val == empty {
		return true
	}

	return false
}

// Any returns true if the predicate is true for any elements in data.
func Any[T any](predicate func(T) bool, data ...T) bool {
	for _, val := range data {
		if predicate(val) {
			return true
		}
	}

	return false
}
