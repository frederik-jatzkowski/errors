// Package format provides functionality to work with format strings.
package format

import (
	"fmt"
)

type String string

func (s String) SplitIntoComponents(args []any) (components Components) {
	prefix, errOut, remainingFormat, remainingArgs := s.ProceedToNextError(s.Runes(), args)
	if prefix != "" {
		components.Components = append(components.Components, prefix)
	}
	if errOut != nil {
		components.Components = append(components.Components, errOut)
		components.Errs = append(components.Errs, errOut)
	}

	for len(remainingFormat) > 0 {
		prefix, errOut, remainingFormat, remainingArgs = s.ProceedToNextError(
			remainingFormat,
			remainingArgs,
		)
		if prefix != "" {
			components.Components = append(components.Components, prefix)
		}

		if errOut != nil {
			components.Components = append(components.Components, errOut)
			components.Errs = append(components.Errs, errOut)
		}
	}

	if len(components.Components) == 0 {
		components.Components = append(components.Components, "")
	}

	return components
}

func (s String) ProceedToNextError(
	format []rune,
	args []any,
) (
	prefix string,
	err error, // nolint: staticcheck
	remainingFormat []rune,
	remainingArgs []any,
) {
	iArg := 0

	currentPartialFormat := ""
	isVerb := false
	i := 0
	for ; i < len(format); i++ {
		r := format[i]

		if r == '%' {
			isVerb = !isVerb
		}

		if isVerb && iArg < len(args) {
			err, ok := args[iArg].(error)
			nextIndex := min(i+1, len(format)-1)
			if format[nextIndex] == 'w' && ok {
				return fmt.Sprintf(currentPartialFormat, args[:iArg]...),
					err,
					format[i+2:],
					args[iArg+1:]
			} else if format[nextIndex] != '%' {
				iArg++
				isVerb = false
			}
		}

		currentPartialFormat += string(r)
	}

	return fmt.Sprintf(currentPartialFormat, args...),
		nil,
		nil,
		nil
}

func (s String) Runes() []rune {
	result := make([]rune, 0, len(s))
	for _, r := range s {
		result = append(result, r)
	}

	return result
}
