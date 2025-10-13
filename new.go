package errors

type simple struct {
	msg string
}

// New works exactly like the standard library [errors.New] but adds a stack trace.
func New(text string) error {
	return innerNew(text)
}

func innerNew(text string) error {
	st := newStackTrace()

	err := &simple{
		msg: text,
	}

	if !st.isSentinel() {
		return ensureStackTrace(err)
	}

	return err
}

func (s *simple) Error() string {
	return s.msg
}
