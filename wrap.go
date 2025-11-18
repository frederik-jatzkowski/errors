package errors

import "errors"

// Is works exactly like the standard library [errors.Is].
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As works exactly like the standard library [errors.As].
func As(err error, target any) bool {
	return errors.As(err, target)
}

// Unwrap works exactly like the standard library [errors.Unwrap].
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
