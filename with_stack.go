package errors

import (
	"fmt"
)

type withStack struct {
	inner error
	st    *stackTrace
}

func (err *withStack) Format(s fmt.State, verb rune) {
	delegateFormat(err.inner, s, verb)
	if shouldPrintStack(s, verb) {
		err.st.Format(s, verb)
	}
}

func (err *withStack) Error() string {
	return err.inner.Error()
}

func (err *withStack) Unwrap() error {
	return err.inner
}
