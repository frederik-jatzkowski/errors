package errors

import "fmt"

// Errorf works exactly like the standard library [fmt.Errorf] but adds a stack trace if none is present.
// It supports the %w verb, including wrapping multiple errors.
//
// If multiple errors are wrapped, it adds a stack trace if at least one of the wrapped errors lacks one.
func Errorf(format string, args ...any) error {
	return innerErrorf(format, args...)
}

func innerErrorf(format string, args ...any) error {
	// Hint for go vet’s printf analyzer to classify this as Errorf-like.
	// nolint: forbidigo
	err := fmt.Errorf(format, args...) // enable printf checking & %w

	unwrapSingle, ok := err.(interface{ Unwrap() error })
	if ok {
		wrapped := unwrapSingle.Unwrap()
		return ensureStackTraceIfNecessary(&errorfSingle{
			format:  format,
			args:    args,
			wrapped: wrapped,
			msg:     err.Error(),
		}, []error{wrapped})
	}

	unwrapMany, ok := err.(interface{ Unwrap() []error })
	if ok {
		wrapped := unwrapMany.Unwrap()
		return ensureStackTraceIfNecessary(&errorfMany{
			format:  format,
			args:    args,
			wrapped: wrapped,
			msg:     err.Error(),
		}, wrapped)
	}

	return ensureStackTraceIfNecessary(&simple{
		msg: err.Error(),
	}, nil)
}
