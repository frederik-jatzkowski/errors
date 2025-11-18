package internal

type Simple struct {
	Stack *WithStack
	Msg   string
}

func NewSimple(depth int, text string) error {
	return EnsureStackTraceIfNecessary(depth+1, &Simple{
		Msg: text,
	}, nil)
}

func (err *Simple) Error() string {
	return err.Msg
}

func (err *Simple) SetWithStack(ws *WithStack) {
	err.Stack = ws
}

func (err *Simple) As(target any) bool {
	switch t := target.(type) {
	case **Simple:
		*t = err
		return true
	case **WithStack:
		if err.Stack == nil {
			return false
		}

		*t = err.Stack
		return true
	}

	return false
}
