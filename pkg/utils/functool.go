package utils

import "github.com/pkg/errors"

func Must[T any](value T, err error) T {
	if err != nil {
		panic(errors.Wrap(err, "call should not failed"))
	}

	return value
}
