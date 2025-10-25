package internal

type ErrorfSingle struct {
	Stack   *WithStack
	format  string
	Msg     string
	Wrapped error
	args    []interface{}
}

func (e *ErrorfSingle) Error() string {
	return e.Msg
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
