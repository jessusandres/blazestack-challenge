package utils

import (
	"blazestack.com/ms-incidents/cmd/interfaces"
)

func HandleServiceError(c interfaces.IAborter, err error) bool {
	if err != nil {
		_ = c.Error(err)
		c.Abort()

		return true
	}

	return false
}

type Result[T any] struct {
	Value T
	Error error
}

func Try[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Error: err}
}

func (r Result[T]) Must() T {
	if r.Error != nil {
		panic(r.Error)
	}
	return r.Value
}
