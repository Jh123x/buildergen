package utils

func Filter[T any](lst []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(lst))

	for _, item := range lst {
		if !predicate(item) {
			continue
		}

		result = append(result, item)
	}

	return result
}

func Map[T, R any](lst []T, mapper func(T) R) []R {
	result := make([]R, 0, len(lst))

	for _, item := range lst {
		result = append(result, mapper(item))
	}

	return result
}

func FilterNil[T any](lst []*T) []*T {
	return Filter(lst, func(v *T) bool { return v != nil })
}
