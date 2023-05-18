package types

// Map is a helper function to map a slice to a new slice of type U.
func Map[T, U any](ts []T, f func(t T, i int) U) []U {
	rs := make([]U, len(ts))
	for i := range ts {
		rs[i] = f(ts[i], i)
	}
	return rs
}

func MapErr[T, U any](ts []T, f func(t T, i int) (U, error)) (rs []U, err error) {
	rs = make([]U, len(ts))
	for i := range ts {
		rs[i], err = f(ts[i], i)
		if err != nil {
			return nil, err
		}
	}
	return rs, nil
}

func Ref[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	return *t
}
