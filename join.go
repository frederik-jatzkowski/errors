package errors

import (
	"fmt"
	"slices"
)

type join struct {
	stack   *withStack
	wrapped []error
}

// Join works exactly like the standard library [errors.Join] but adds a stack trace if at least one of the wrapped errors
// lacks one.
func Join(errs ...error) error {
	return innerJoin(errs...)
}

func innerJoin(errs ...error) error {
	errs = slices.DeleteFunc(errs, func(err error) bool {
		return err == nil
	})

	if len(errs) == 0 {
		return nil
	}

	return ensureStackTraceIfNecessary(&join{
		wrapped: errs,
	}, errs)
}

func (j *join) Error() string {
	return fmt.Sprint(j)
}

func (j *join) Unwrap() []error {
	return j.wrapped
}

func (j *join) setWithStack(ws *withStack) {
	j.stack = ws
}

func (j *join) As(target any) bool {
	switch t := target.(type) {
	case **join:
		*t = j
		return true
	case **withStack:
		if j.stack == nil {
			return false
		}

		*t = j.stack
		return true
	}

	return false
}
