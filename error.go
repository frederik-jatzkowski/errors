package errors

func ensureStackTrace(err error) error {
	trace := &withStack{}
	if As(err, &trace) {
		return err
	}

	return &withStack{
		inner: err,
		st:    newStackTrace(),
	}
}
