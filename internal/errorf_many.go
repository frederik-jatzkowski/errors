package internal

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors/internal/format"
)

type ErrorfMany struct {
	Components format.Components
	Stack      *WithStack
}

func (e *ErrorfMany) Error() string {
	return fmt.Sprint(e)
}

func (e *ErrorfMany) SetWithStack(ws *WithStack) {
	e.Stack = ws
}

func (e *ErrorfMany) As(target any) bool {
	switch t := target.(type) {
	case **ErrorfMany:
		*t = e
		return true
	case **WithStack:
		if e.Stack == nil {
			return false
		}

		*t = e.Stack
		return true
	}

	return false
}

func (e *ErrorfMany) Unwrap() []error {
	return e.Components.Errs
}
