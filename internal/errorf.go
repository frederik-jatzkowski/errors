package internal

import "fmt"

func Errorf(depth int, format string, args ...any) error {
	// Hint for the govet printf analyzer to classify this as Errorf-like.
	// nolint: forbidigo
	err := fmt.Errorf(format, args...) // enable printf checking & %w

	unwrapSingle, ok := err.(interface{ Unwrap() error })
	if ok {
		wrapped := unwrapSingle.Unwrap()
		// nolint: wrapcheck
		return EnsureStackTraceIfNecessary(depth+1, &ErrorfSingle{
			format:  format,
			args:    args,
			Wrapped: wrapped,
			Msg:     err.Error(),
		}, []error{wrapped})
	}

	unwrapMany, ok := err.(interface{ Unwrap() []error })
	if ok {
		wrapped := unwrapMany.Unwrap()
		// nolint: wrapcheck
		return EnsureStackTraceIfNecessary(depth+1, &ErrorfMany{
			format:  format,
			args:    args,
			Wrapped: wrapped,
			Msg:     err.Error(),
		}, wrapped)
	}

	// nolint: wrapcheck
	return EnsureStackTraceIfNecessary(depth+1, &Simple{
		Msg: err.Error(),
	}, nil)
}
