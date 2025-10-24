package errors

import (
	"fmt"
	"runtime"
	"strings"
)

type stackTrace runtime.StackRecord

func newStackTrace() *stackTrace {
	record := &stackTrace{}
	runtime.Callers(5, record.Stack0[:])
	return record
}

func (st *stackTrace) String() string {
	return fmt.Sprintf("%+v", st)
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

// ensureStackTraceIfNecessary adds a stack trace to the error if at least one of the
// shouldHaveStack errors doesn't have one already.
// It also sets the first instance of *withStack found in err.
func ensureStackTraceIfNecessary(err internalError, shouldHaveStack []error) error {
	var (
		trace                *withStack
		actuallyHavingStacks int
	)
	for _, err2 := range shouldHaveStack {
		if As(err2, &trace) {
			if actuallyHavingStacks == 0 {
				err.setWithStack(trace)
			}
			actuallyHavingStacks++
		}
	}

	if actuallyHavingStacks == len(shouldHaveStack) && len(shouldHaveStack) > 0 {
		return err
	}

	return &withStack{
		inner: err,
		st:    newStackTrace(),
	}
}
