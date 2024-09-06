package utils

import "unicode"

// IsNilOrEmpty returns true if the value is nil or equal to the empty value.
func IsNilOrEmpty(val *string) bool {
	return val == nil || len(*val) == 0
}

func LowerFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToLower(r[0])

	return string(r)
}
