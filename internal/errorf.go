package internal

import (
	"fmt"

	internalFormat "github.com/frederik-jatzkowski/errors/internal/format"
)

func Errorf(depth int, format string, args ...any) error {
	// Hint for the govet printf analyzer to classify this as Errorf-like.

	err := fmt.Errorf(format, args...) // enable printf checking & %w

	components := internalFormat.String(format).SplitIntoComponents(args)
	if len(components.Errs) > 1 {
		return EnsureStackTraceIfNecessary(depth+1, &ErrorfMany{
			Components: components,
		}, components.Errs)
	}

	if len(components.Errs) == 1 {
		return EnsureStackTraceIfNecessary(depth+1, &ErrorfSingle{
			components: components,
			Wrapped:    components.Errs[0],
		}, components.Errs)
	}

	return EnsureStackTraceIfNecessary(depth+1, &Simple{
		Msg: err.Error(),
	}, nil)
}
