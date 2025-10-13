package errors

import "errors"

// Is works exactly like the standard library [errors.Is].
func Is(err, target error) bool {
	return errors.Is(err, target) // nolint: forbidigo
}

// As works exactly like the standard library [errors.As].
func As(err error, target any) bool {
	return errors.As(err, target) // nolint: forbidigo
}

// Unwrap works exactly like the standard library [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err) // nolint: forbidigo
}
