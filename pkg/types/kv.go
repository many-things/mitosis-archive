package types

// KV is base data structure to store key-value set effectively.
type KV[K, V any] struct {
	Key   K
	Value V
}

func NewKV[K, V any](key K, value V) KV[K, V] {
	return KV[K, V]{Key: key, Value: value}
}

func Keys[K, V any](kvs []KV[K, V]) []K {
	ks := make([]K, len(kvs))
	for i, kv := range kvs {
		ks[i] = kv.Key
	}
	return ks
}

func Values[K, V any](kvs []KV[K, V]) []V {
	vs := make([]V, len(kvs))
	for i, kv := range kvs {
		vs[i] = kv.Value
	}
	return vs
}

// MapKV is a helper function to map a KV slice to a new slice of type R.
func MapKV[K, V, R any](kvs []KV[K, V], f func(k K, v V) R) []R {
	rs := make([]R, len(kvs))
	for i := range kvs {
		kv := kvs[i]
		rs[i] = f(kv.Key, kv.Value)
	}
	return rs
}
