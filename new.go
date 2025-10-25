package errors

import "github.com/frederik-jatzkowski/errors/internal"

// New works exactly like the standard library [errors.New] but adds a stack trace.
//
// If it is called from a sentinel context, no stack trace is added.
func New(text string) error {
	return internal.NewSimple(1, text) // nolint: wrapcheck
}
