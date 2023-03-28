package utils

import "github.com/pkg/errors"

func Must[T any](value T, err error) T {
	if err != nil {
		panic(errors.Wrap(err, "call should not failed"))
	}

	return value
}

func Any[T any](source []T, checker func(T) bool) bool {
	for _, val := range source {
		if checker(val) {
			return true
		}
	}

	return false
}

func All[T any](source []T, checker func(T) bool) bool {
	for _, val := range source {
		if !checker(val) {
			return false
		}
	}
	return true
}
