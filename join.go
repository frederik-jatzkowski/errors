package errors

import (
	"fmt"
)

type join struct {
	wrapped []error
}

// Join works exactly like the standard library [errors.Join] but adds a stack trace if none is present.
func Join(errs ...error) error {
	return innerJoin(errs...)
}

func innerJoin(errs ...error) error {
	var nonNil []error
	for _, err := range errs {
		if err != nil {
			nonNil = append(nonNil, err)
		}
	}

	if len(nonNil) == 0 {
		return nil
	}

	return ensureStackTrace(&join{
		wrapped: nonNil,
	})
}

func (j *join) Error() string {
	return fmt.Sprint(j)
}

func (j *join) Format(f fmt.State, verb rune) {
	for _, err := range j.wrapped {
		delegateFormat(err, f, verb)
		// nolint: errcheck
		fmt.Fprintln(f)
	}
}

func (j *join) Unwrap() []error {
	return j.wrapped
}
