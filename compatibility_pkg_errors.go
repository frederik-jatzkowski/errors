package errors

import "github.com/frederik-jatzkowski/errors/internal"

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement the Cause() method, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
//
// This function allows compatibility for [github.com/pkg/errors.Cause] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Use [Is]
// and [As] for error inspection.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if causer, ok := err.(interface{ Cause() error }); ok {
		return causer.Cause()
	}

	return err
}

// WithMessage annotates err with a new message.
// If err is nil, [WithMessage] returns nil.
//
// This implementation uses the module's [Errorf] functionality internally,
// which means it will add stack traces unless the error already has one.
//
// This function allows compatibility for [github.com/pkg/errors.WithMessage] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Use [Errorf]
// with %w verb for wrapping errors with messages.
func WithMessage(err error, msg string) error {
	if err == nil {
		return nil
	}

	return internal.Errorf(1, "%s: %w", msg, err)
}

// WithMessagef annotates err with the format specifier and arguments.
// If err is nil, [WithMessagef] returns nil.
//
// This implementation uses the module's [Errorf] functionality internally,
// which means it will add stack traces unless the error already has one.
//
// This function allows compatibility for [github.com/pkg/errors.WithMessagef] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Use [Errorf]
// with %w verb for wrapping errors with formatted messages.
func WithMessagef(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return internal.Errorf(1, format+": %w", append(args, err)...)
}

// WithStack annotates err with a stack trace at the point [WithStack] was called.
// If err is nil, [WithStack] returns nil.
//
// This implementation uses the module's [Join] functionality internally to ensure
// stack trace information is preserved and added appropriately.
//
// This function allows compatibility for [github.com/pkg/errors.WithStack] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Stack traces are
// automatically added by [New], [Errorf], and [Join] functions.
func WithStack(err error) error {
	return internal.NewJoin(1, err)
}

// Wrap returns an error annotating err with a stack trace at the point [Wrap] is
// called, and the supplied message. If err is nil, [Wrap] returns nil.
//
// This implementation uses the module's [Errorf] functionality internally,
// which means it will add stack traces unless the error already has one.
//
// This function allows compatibility for [github.com/pkg/errors.Wrap] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Use [Errorf]
// with %w verb: [Errorf]("message: %w", err).
func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	return internal.Errorf(1, "%s: %w", msg, err)
}

// Wrapf returns an error annotating err with a stack trace at the point [Wrapf] is
// called, and the format specifier. If err is nil, [Wrapf] returns nil.
//
// This implementation uses the module's [Errorf] functionality internally,
// which means it will add stack traces unless the error already has one.
//
// This function allows compatibility for [github.com/pkg/errors.Wrapf] but is
// currently not encouraged by this module.
//
// Deprecated: Use the recommended API of this module instead. Use [Errorf]
// with %w verb for wrapping errors with formatted messages.
func Wrapf(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return internal.Errorf(1, format+": %w", append(args, err)...)
}
