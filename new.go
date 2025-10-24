package errors

type simple struct {
	stack *withStack
	msg   string
}

// New works exactly like the standard library [errors.New] but adds a stack trace.
//
// If it is called from a sentinel context, no stack trace is added.
func New(text string) error {
	return innerNew(text)
}

func innerNew(text string) error {
	st := newStackTrace()

	err := &simple{
		msg: text,
	}

	if !st.isSentinel() {
		return ensureStackTraceIfNecessary(err, nil)
	}

	return err
}

func (s *simple) Error() string {
	return s.msg
}

func (s *simple) setWithStack(ws *withStack) {
	s.stack = ws
}

func (s *simple) As(target any) bool {
	switch t := target.(type) {
	case **simple:
		*t = s
		return true
	case **withStack:
		if s.stack == nil {
			return false
		}

		*t = s.stack
		return true
	}

	return false
}
