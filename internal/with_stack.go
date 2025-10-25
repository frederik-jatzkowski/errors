package internal

import (
	"fmt"
)

type WithStack struct {
	Inner error
	St    *StackTrace
}

func (err *WithStack) Format(s fmt.State, verb rune) {
	delegateFormat(err.Inner, s, verb)
	if shouldPrintStack(s, verb) {
		err.St.Format(s, verb)
	}
}

func (err *WithStack) Error() string {
	return err.Inner.Error()
}

func (err *WithStack) Unwrap() error {
	return err.Inner
}
