package errors

type errorfSingle struct {
	stack   *withStack
	format  string
	msg     string
	wrapped error
	args    []interface{}
}

func (e *errorfSingle) Error() string {
	return e.msg
}

func (e *errorfSingle) Unwrap() error {
	return e.wrapped
}

func (e *errorfSingle) setWithStack(ws *withStack) {
	e.stack = ws
}

func (e *errorfSingle) As(target any) bool {
	switch t := target.(type) {
	case **errorfSingle:
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
