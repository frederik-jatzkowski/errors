package errors

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

func getStackTrace(err error) (st *stackTrace, ok bool) {
	var ws *withStack
	if As(err, &ws) {
		return ws.st, true
	}

	return nil, false
}

// SprintStackTrace extracts and formats the stack trace from an error as a string.
// If the error does not contain stack trace information, it returns an empty string.
func SprintStackTrace(err error) string {
	st, ok := getStackTrace(err)
	if !ok {
		return ""
	}

	return st.String()
}

type stackTrace runtime.StackRecord

func newStackTrace() *stackTrace {
	record := &stackTrace{}
	runtime.Callers(5, record.Stack0[:])
	return record
}

func (st *stackTrace) String() string {
	return fmt.Sprintf("%+v\n", st)
}

func (st *stackTrace) Format(s fmt.State, verb rune) {
	switch {
	case s.Flag('#'):
		// nolint: errcheck
		fmt.Fprintf(s, "%#v", (*runtime.StackRecord)(st))
	default:
		for _, pc := range st.Stack0 {
			if pc == 0 {
				continue
			}

			// nolint: errcheck
			io.WriteString(s, "\n")

			// nolint: errcheck
			s.Write([]byte(formatFrame(pc)))
		}
	}
}

func (st *stackTrace) isSentinel() bool {
	for _, pc := range st.Stack0 {
		if pc == 0 {
			continue
		}
		name := runtime.FuncForPC(pc).Name()
		if !strings.HasPrefix(name, "runtime.") && !strings.HasSuffix(name, ".init") {
			return false
		}
	}

	return true
}

func formatFrame(pc uintptr) string {
	funcForPC := runtime.FuncForPC(pc)
	if funcForPC == nil {
		return "unknown"
	}

	name := funcForPC.Name()
	file, line := funcForPC.FileLine(pc)

	return fmt.Sprintf("\t%s\n\t\t%s:%d", name, file, line)
}

func (st *stackTrace) RuntimeStackRecord() *runtime.StackRecord {
	return (*runtime.StackRecord)(st)
}
