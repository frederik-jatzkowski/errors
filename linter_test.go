package errors_test

import (
	"fmt"
	"io"
	"strconv"
	"testing"

	"github.com/frederik-jatzkowski/errors"
)

func ErrA() error {
	_, err := strconv.ParseInt("a", 10, 64)
	return err // nolint: wrapcheck
}

func ErrB() error {
	_, err := strconv.ParseInt("b", 10, 64)
	return errors.Join(err)
}

// TestWrapcheck checks if the linter actually complains when configured as described in our documentation.
func TestWrapcheck(t *testing.T) {
	// nolint: errcheck
	fmt.Fprintf(io.Discard, "%+v\n", ErrA())
	// nolint: errcheck
	fmt.Fprintf(io.Discard, "%+v\n", ErrB())
}
