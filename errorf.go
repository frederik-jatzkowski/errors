package errors

import (
	"github.com/frederik-jatzkowski/errors/internal"
)

// Errorf works exactly like the standard library [fmt.Errorf] but adds a stack trace if none is present.
// It supports the %w verb, including wrapping multiple errors.
//
// If multiple errors are wrapped, it adds a stack trace if at least one of the wrapped errors lacks one.
func Errorf(format string, args ...any) error {
	return internal.Errorf(1, format, args...)
}
