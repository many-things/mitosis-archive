package types

// Map is a helper function to map a slice to a new slice of type U.
func Map[T, U any](ts []T, f func(t T) U) []U {
	rs := make([]U, len(ts))
	for i := range ts {
		rs[i] = f(ts[i])
	}
	return rs
}

func Ref[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	return *t
}
