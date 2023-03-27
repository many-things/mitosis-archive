package utils

func ForEach[T any](source []T, f func(T)) {
	for i := range source {
		f(source[i])
	}
}
