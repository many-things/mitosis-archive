package queue

func min[T ~int | ~int64 | ~uint | ~uint64](x, y T) T {
	if x > y {
		return y
	} else {
		return x
	}
}
