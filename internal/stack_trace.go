package internal

import (
	"errors"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type StackTrace runtime.StackRecord

func NewStackTrace(depth int) *StackTrace {
	record := &StackTrace{}
	runtime.Callers(depth+2, record.Stack0[:])
	return record
}

func (st *StackTrace) String() string {
	return fmt.Sprintf("%+v", st)
}

func (st *StackTrace) IsSentinel() bool {
	if st.Stack0[0] == 0 {
		return true
	}

	firstFuncName := runtime.FuncForPC(st.Stack0[0]).Name()
	firstFuncBase := filepath.Base(firstFuncName)

	return strings.HasSuffix(firstFuncBase, ".init")
}

// EnsureStackTraceIfNecessary adds a stack trace to the error if at least one of the
// shouldHaveStack errors doesn't have one already.
// It also sets the first instance of *withStack found in err.
func EnsureStackTraceIfNecessary(depth int, err Error, shouldHaveStack []error) error {
	var (
		trace                *WithStack
		actuallyHavingStacks int
	)
	for _, err2 := range shouldHaveStack {
		if errors.As(err2, &trace) {
			if actuallyHavingStacks == 0 {
				err.SetWithStack(trace)
			}
			actuallyHavingStacks++
		}
	}

	if actuallyHavingStacks == len(shouldHaveStack) && len(shouldHaveStack) > 0 {
		return err
	}

	st := NewStackTrace(depth + 1)
	if st.IsSentinel() {
		return err
	}

	return &WithStack{
		Inner: err,
		St:    st,
	}
}
