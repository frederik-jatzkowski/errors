package internal

import (
	"fmt"
	"slices"
)

type Join struct {
	Stack   *WithStack
	Wrapped []error
}

func NewJoin(depth int, errs ...error) error {
	errs = slices.DeleteFunc(errs, func(err error) bool {
		return err == nil
	})

	if len(errs) == 0 {
		return nil
	}

	return EnsureStackTraceIfNecessary(depth+1, &Join{
		Wrapped: errs,
	}, errs)
}

func (err *Join) Error() string {
	return fmt.Sprint(err)
}

func (err *Join) Unwrap() []error {
	return err.Wrapped
}

func (err *Join) SetWithStack(ws *WithStack) {
	err.Stack = ws
}

func (err *Join) As(target any) bool {
	switch t := target.(type) {
	case **Join:
		*t = err
		return true
	case **WithStack:
		if err.Stack == nil {
			return false
		}

		*t = err.Stack
		return true
	}

	return false
}
