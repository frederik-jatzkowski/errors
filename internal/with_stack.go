package internal

type WithStack struct {
	Inner    error
	St       *StackTrace
	Explicit bool
}

func (e *WithStack) Error() string {
	return e.Inner.Error()
}

func (e *WithStack) Unwrap() error {
	return e.Inner
}
