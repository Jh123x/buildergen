package utils

func IsNilOrEmpty[inputType comparable](val *inputType) bool {
	var empty inputType
	if val == nil || *val == empty {
		return true
	}

	return false
}

func Any[T any](predicate func(T) bool, data ...T) bool {
	for _, val := range data {
		if predicate(val) {
			return true
		}
	}

	return false
}
