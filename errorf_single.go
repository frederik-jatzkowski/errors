package errors

import (
	"fmt"
)

type errorfSingle struct {
	wrapped error
	msg     string
}

func (e *errorfSingle) Error() string {
	return e.msg
}

func (e *errorfSingle) Format(f fmt.State, verb rune) {
	printErrorString(e.msg, f, verb)
	if shouldPrintStack(f, verb) {
		// nolint: errcheck
		fmt.Fprintln(f)
		delegateFormat(e.wrapped, f, verb)
	}
}

func (e *errorfSingle) Unwrap() error {
	return e.wrapped
}
