package errors_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

var (
	ErrTest = errors.New("test error")
)

func Example() {
	err := errors.Errorf("an error occurred: %w", ErrTest)
	fmt.Println(err)
	fmt.Println("can check for wrapped sentinel errors:", errors.Is(err, ErrTest))
	fmt.Println("sentinel error has no stack trace:", errors.SprintStackTrace(ErrTest) == "")
	fmt.Println("wrapped error has a stack trace:", errors.SprintStackTrace(err) != "")

	extraErr := errors.Errorf("extra: %w", err)
	stackTrace := errors.SprintStackTrace(extraErr)

	errFmt := fmt.Sprintf("%+v", err)
	extraErrFmt := fmt.Sprintf("%+v", extraErr)

	fmt.Println("prevents duplicate stack traces:", len(extraErrFmt) < len(errFmt)+len(stackTrace))

	// Output:
	// an error occurred: test error
	// can check for wrapped sentinel errors: true
	// sentinel error has no stack trace: true
	// wrapped error has a stack trace: true
	// prevents duplicate stack traces: true
}
