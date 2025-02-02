package utils

func Filter[T any](lst []T, predicate func(T) bool) []T {
	if predicate == nil {
		return lst
	}

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
	acc := make([]*T, 0, len(lst))
	for _, v := range lst {
		if v == nil {
			continue
		}

		acc = append(acc, v)
	}

	return acc
}
