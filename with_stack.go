package errors

import "github.com/frederik-jatzkowski/errors/internal"

// WithStack annotates err with a stack trace at the point [WithStack] was called, even if one is already present.
// If err is nil, [WithStack] returns nil.
//
// This allows explicit recording of stack traces when needed.
// The most likely use case for this is wrapping errors that originated in other goroutines.
// Because each goroutine has its own stack, you might want to capture additionally stack traces when the error tree spans multiple goroutines.
//
// Do not use this under normal operations, use [Errorf] and [Join] for wrapping errors.
func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return &internal.WithStack{
		Explicit: true,
		Inner:    err,
		St:       internal.NewStackTrace(1),
	}
}
