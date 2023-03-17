package types

// Map is a helper function to map a slice to a new slice of type U.
func Map[K, V, R any](kvs []KV[K, V], f func(k K, v V) R) []R {
	rs := make([]R, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		rs[i] = f(kv.Key, kv.Value)
	}
	return rs
}

func Ref[T any](t T) *T {
	return &t
}

func Deref[T any](t *T) T {
	return *t
}
