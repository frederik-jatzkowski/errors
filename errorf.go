package errors

import "fmt"

// Errorf formats according to a format specifier and returns the string as a
// value that satisfies error. Errorf also records the stack trace at the point
// it was called.
//
// If the format specifier includes a %w verb with an error operand, the returned
// error will implement an Unwrap method returning the operand. If there is more
// than one %w verb, the returned error will implement an Unwrap method returning
// a slice of all error operands. It is invalid to include an operand for a %w verb
// that does not implement the error interface.
// The %w verb is otherwise a synonym for %v.
//
// This function is compatible with Go's standard fmt.Errorf but adds automatic
// stack trace collection.
//
// Example:
//
//	base := errors.New("base error")
//	wrapped := errors.Errorf("wrapped: %w", base)
//	fmt.Printf("%+v", wrapped)  // prints error with stack trace
func Errorf(format string, args ...any) error {
	return innerErrorf(format, args...)
}

func innerErrorf(format string, args ...any) error {
	// Hint for go vet’s printf analyzer to classify this as Errorf-like.
	// nolint: forbidigo
	msg := fmt.Errorf(format, args...) // enable printf checking & %w

	var errs []error

	firstPercent := false
	verbIndex := 0
	for _, char := range format {
		if char == ' ' {
			continue
		}

		if char == '%' {
			if firstPercent {
				firstPercent = false
				continue
			}

			firstPercent = true
			continue
		}

		if !firstPercent {
			continue
		}

		if verbIndex >= len(args) {
			break
		}

		err, ok := args[verbIndex].(error)
		if ok && char == 'w' {
			errs = append(errs, err)
		}

		verbIndex++
		firstPercent = false
	}

	if len(errs) == 1 {
		return ensureStackTrace(&errorfSingle{
			msg:     msg.Error(),
			wrapped: errs[0],
		})
	}

	return ensureStackTrace(&errorfMany{
		msg:     msg.Error(),
		wrapped: errs,
	})
}
