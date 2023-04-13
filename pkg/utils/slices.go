package utils

func ForEach[T any](source []T, f func(T)) {
	for i := range source {
		f(source[i])
	}
}

func IndexOf[T comparable](target T, elems []T) int {
	for i, v := range elems {
		if v == target {
			return i
		}
	}

	return -1
}
