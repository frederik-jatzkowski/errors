package errors

import (
	"fmt"
)

type errorfMany struct {
	msg     string
	wrapped []error
}

func (e *errorfMany) Error() string {
	return e.msg
}

func (e *errorfMany) Format(f fmt.State, verb rune) {
	printErrorString(e.msg, f, verb)
	if shouldPrintStack(f, verb) && len(e.wrapped) > 0 {
		// nolint: errcheck
		fmt.Fprintln(f)
		for _, err := range e.wrapped {
			delegateFormat(err, f, verb)
			// nolint: errcheck
			fmt.Fprintln(f)
		}
	}
}
