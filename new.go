package errors

type simple struct {
	msg string
}

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
//
// This function automatically adds stack trace information to the error unless
// it's called during package initialization or other sentinel contexts.
//
// Example:
//
//	err := errors.New("something went wrong")
//	fmt.Printf("%+v", err)  // prints error with stack trace
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
