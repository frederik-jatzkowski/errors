package internal

import (
	"fmt"
	"runtime"

	"github.com/frederik-jatzkowski/errors/internal/dto"
)

func (e *ErrorfMany) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		_ = e.ToDTO(nil).WriteLong(dto.NewWriter(s, 0))
	case shouldPrintErrorMessage(s, verb):
		_ = e.ToDTO(nil).WriteShort(dto.NewWriter(s, 0))
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (e *ErrorfSingle) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb):
		_ = e.ToDTO(nil).WriteLong(dto.NewWriter(s, 0))
	case shouldPrintErrorMessage(s, verb):
		_ = e.ToDTO(nil).WriteShort(dto.NewWriter(s, 0))
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%v)", verb, e)
	}
}

func (err *Join) Format(s fmt.State, verb rune) {
	if shouldPrintStack(s, verb) {
		_ = err.ToDTO(nil).WriteLong(dto.NewWriter(s, -1))
	} else {
		_ = err.ToDTO(nil).WriteShort(dto.NewWriter(s, -1))
	}
}

func (err *WithStack) Format(s fmt.State, verb rune) {
	if shouldPrintStack(s, verb) {
		// nolint: errcheck
		fmt.Fprintf(s, fmt.FormatString(s, verb), err.Inner)
		// nolint: errcheck
		fmt.Fprintf(s, fmt.FormatString(s, verb), err.St)
	} else {
		// nolint: errcheck
		fmt.Fprintf(s, fmt.FormatString(s, verb), err.Inner)
	}
}

func shouldPrintStack(s fmt.State, verb rune) bool {
	return verb == 'v' && s.Flag('+')
}

func shouldPrintErrorMessage(s fmt.State, verb rune) bool {
	return verb == 'v' && !s.Flag('+') || verb == 's'
}

func (st *StackTrace) Format(s fmt.State, verb rune) {
	switch {
	case shouldPrintStack(s, verb) || verb == 's':
		for _, pc := range st.Stack0 {
			if pc == 0 {
				break
			}

			// nolint: errcheck
			fmt.Fprintln(s)

			// nolint: errcheck
			s.Write([]byte(FormatFrame(pc)))
		}
	default:
		// nolint: errcheck
		fmt.Fprintf(s, "%%!%c(%s)", verb, st)
	}

	// nolint: errcheck
	fmt.Fprintln(s)
}

func FormatFrame(pc uintptr) string {
	funcForPC := runtime.FuncForPC(pc)
	if funcForPC == nil {
		return "unknown"
	}

	name := funcForPC.Name()
	file, line := funcForPC.FileLine(pc)

	return fmt.Sprintf("    %s\n        %s:%d", name, file, line)
}
