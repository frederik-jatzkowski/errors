package errors

import (
	"fmt"
)

type join struct {
	wrapped []error
}

// Join returns an error that wraps the given errors.
// Any nil error values are discarded.
// Join returns nil if all errors are nil.
// The error formats as the concatenation of the strings obtained
// by calling the Error method of each element of errs, with a newline
// between each string.
//
// A non-nil error returned by Join implements the Unwrap() []error method,
// which returns a copy of errs with any nil values removed.
//
// This function automatically adds stack trace information to the joined error.
//
// Example:
//
//	err1 := errors.New("first error")
//	err2 := errors.New("second error")
//	joined := errors.Join(err1, err2)
//	fmt.Printf("%+v", joined)  // prints both errors with stack trace
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
