package utils

// IsNilOrEmpty returns true if the value is nil or equal to the empty value.
func IsNilOrEmpty(val *string) bool {
	return val == nil || len(*val) == 0
}
