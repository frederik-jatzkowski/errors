package internal

type WithStack struct {
	Inner error
	St    *StackTrace
}

func (err *WithStack) Error() string {
	return err.Inner.Error()
}

func (err *WithStack) Unwrap() error {
	return err.Inner
}
