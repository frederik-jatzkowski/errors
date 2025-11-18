package internal

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors/internal/format"
)

type ErrorfSingle struct {
	Wrapped    error
	Stack      *WithStack
	components format.Components
}

func (e *ErrorfSingle) Error() string {
	return fmt.Sprint(e)
}

func (e *ErrorfSingle) Unwrap() error {
	return e.Wrapped
}

func (e *ErrorfSingle) SetWithStack(ws *WithStack) {
	e.Stack = ws
}

func (e *ErrorfSingle) As(target any) bool {
	switch t := target.(type) {
	case **ErrorfSingle:
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
