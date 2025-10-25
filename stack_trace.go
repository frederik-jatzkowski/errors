package errors

import (
	"fmt"
	"path/filepath"
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
	if st.Stack0[0] == 0 {
		return true
	}

	firstFuncName := runtime.FuncForPC(st.Stack0[0]).Name()
	firstFuncBase := filepath.Base(firstFuncName)

	return strings.HasSuffix(firstFuncBase, ".init")
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

	st := newStackTrace()
	if st.isSentinel() {
		return err
	}

	return &withStack{
		inner: err,
		st:    st,
	}
}
