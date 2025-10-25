package internal

type Simple struct {
	Stack *WithStack
	Msg   string
}

func NewSimple(depth int, text string) error {
	// nolint: wrapcheck
	return EnsureStackTraceIfNecessary(depth+1, &Simple{
		Msg: text,
	}, nil)
}

func (s *Simple) Error() string {
	return s.Msg
}

func (s *Simple) SetWithStack(ws *WithStack) {
	s.Stack = ws
}

func (s *Simple) As(target any) bool {
	switch t := target.(type) {
	case **Simple:
		*t = s
		return true
	case **WithStack:
		if s.Stack == nil {
			return false
		}

		*t = s.Stack
		return true
	}

	return false
}
