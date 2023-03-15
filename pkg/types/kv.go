package types

type KV[K, V any] struct {
	Key   K
	Value V
}

func NewKV[K, V any](key K, value V) KV[K, V] {
	return KV[K, V]{Key: key, Value: value}
}
