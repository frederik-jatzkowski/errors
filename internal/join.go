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

	// nolint: wrapcheck
	return EnsureStackTraceIfNecessary(depth+1, &Join{
		Wrapped: errs,
	}, errs)
}

func (j *Join) Error() string {
	return fmt.Sprint(j)
}

func (j *Join) Unwrap() []error {
	return j.Wrapped
}

func (j *Join) SetWithStack(ws *WithStack) {
	j.Stack = ws
}

func (j *Join) As(target any) bool {
	switch t := target.(type) {
	case **Join:
		*t = j
		return true
	case **WithStack:
		if j.Stack == nil {
			return false
		}

		*t = j.Stack
		return true
	}

	return false
}
