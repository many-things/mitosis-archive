package utils

import "bytes"

func Must[expected any](e expected, err error) expected {
	if err != nil {
		panic(err.Error())
	}
	return e
}

func Unwrap[expected any](f func() (expected, error)) expected {
	return Must(f())
}
func Unwrap1[expected any, arg1 any](f func(arg1) (expected, error), a1 arg1) expected {
	return Must(f(a1))
}
func Unwrap2[expected any, arg1 any, arg2 any](f func(arg1, arg2) (expected, error), a1 arg1, a2 arg2) expected {
	return Must(f(a1, a2))
}

func JoinBytes(bzs ...[]byte) []byte {
	return bytes.Join(bzs, []byte{})
}
