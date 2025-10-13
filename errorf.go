package errors

import "fmt"

// Errorf works exactly like the standard library [fmt.Errorf] but adds a stack trace if none is present.
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
