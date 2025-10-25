package internal

type ErrorfMany struct {
	format  string
	args    []interface{}
	Msg     string
	Stack   *WithStack
	Wrapped []error
}

func (e *ErrorfMany) Error() string {
	return e.Msg
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
	return e.Wrapped
}
