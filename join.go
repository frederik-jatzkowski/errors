package errors

import (
	"github.com/frederik-jatzkowski/errors/internal"
)

// Join works exactly like the standard library [errors.Join] but adds a stack trace if at least one of the wrapped errors
// lacks one.
func Join(errs ...error) error {
	return internal.NewJoin(1, errs...)
}
