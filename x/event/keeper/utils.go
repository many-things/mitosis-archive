package keeper

import "bytes"

func must[expected any](e expected, err error) expected {
	if err != nil {
		panic(err.Error())
	}
	return e
}

func unwrap[expected any](f func() (expected, error)) expected {
	return must(f())
}
func unwrap1[expected any, arg1 any](f func(arg1) (expected, error), a1 arg1) expected {
	return must(f(a1))
}
func unwrap2[expected any, arg1 any, arg2 any](f func(arg1, arg2) (expected, error), a1 arg1, a2 arg2) expected {
	return must(f(a1, a2))
}

func joinBytes(j []byte, bzs ...[]byte) []byte {
	return bytes.Join(bzs, j)
}
