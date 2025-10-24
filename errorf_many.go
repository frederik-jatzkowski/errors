package errors

type errorfMany struct {
	format  string
	args    []interface{}
	msg     string
	stack   *withStack
	wrapped []error
}

func (e *errorfMany) Error() string {
	return e.msg
}

func (e *errorfMany) setWithStack(ws *withStack) {
	e.stack = ws
}

func (e *errorfMany) As(target any) bool {
	switch t := target.(type) {
	case **errorfMany:
		*t = e
		return true
	case **withStack:
		if e.stack == nil {
			return false
		}

		*t = e.stack
		return true
	}

	return false
}

func (e *errorfMany) Unwrap() []error {
	return e.wrapped
}
