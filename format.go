package errors

import (
	"fmt"
	"runtime"
)

func (e *errorfMany) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		e.formatWithStack(s, verb)
	case shouldPrintErrorMessage(s, verb):
		// nolint: errcheck
		fmt.Fprintf(s, fmt.FormatString(s, verb), e.msg)
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (e *errorfMany) formatWithStack(f fmt.State, verb rune) {
	// nolint: errcheck
	fmt.Fprintf(f, rewriteFormatString(e.format), e.args...)
}

func (e *errorfSingle) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		e.formatWithStack(s, verb)
	case shouldPrintErrorMessage(s, verb):
		// nolint: errcheck
		fmt.Fprintf(s, fmt.FormatString(s, verb), e.msg)
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (e *errorfSingle) formatWithStack(f fmt.State, verb rune) {
	// nolint: errcheck
	fmt.Fprintf(f, rewriteFormatString(e.format), e.args...)
}

func (j *join) Format(f fmt.State, verb rune) {
	if len(j.wrapped) > 1 {
		// nolint: errcheck
		fmt.Fprintln(f)
	}
	for _, err := range j.wrapped {
		delegateFormat(err, f, verb)

		// nolint: errcheck
		fmt.Fprintln(f)
	}
}

func rewriteFormatString(format string) string {
	rewritten := ""
	isVerb := false
	for _, r := range format {
		switch {
		case r == '%' && !isVerb:
			isVerb = true
		case r == 'w' && isVerb:
			isVerb = false
			rewritten += "+v"
			continue
		default:
			isVerb = false

		}

		rewritten += string(r)
	}

	return rewritten
}

func shouldPrintStack(s fmt.State, verb rune) bool {
	return verb == 'v' && s.Flag('+')
}

func shouldPrintErrorMessage(s fmt.State, verb rune) bool {
	return verb == 'v' && !s.Flag('+') || verb == 's'
}

func delegateFormat(err error, s fmt.State, verb rune) {
	// nolint: errcheck
	fmt.Fprintf(s, fmt.FormatString(s, verb), err)
}

func (st *stackTrace) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb) || verb == 's':
		for _, pc := range st.Stack0 {
			if pc == 0 {
				break
			}

			// nolint: errcheck
			fmt.Fprintln(s)

			// nolint: errcheck
			s.Write([]byte(formatFrame(pc)))
		}
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%s)", verb, st)
	}

	// nolint: errcheck
	fmt.Fprintln(s)
}

func formatFrame(pc uintptr) string {
	funcForPC := runtime.FuncForPC(pc)
	if funcForPC == nil {
		return "unknown"
	}

	name := funcForPC.Name()
	file, line := funcForPC.FileLine(pc)

	return fmt.Sprintf("    %s\n        %s:%d", name, file, line)
}
