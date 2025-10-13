package errors_test

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/frederik-jatzkowski/errors"
)

func Example_ignoreStackTrace() {
	err := errors.New("hello world")
	fmt.Println(err)
	// Output:
	// hello world
}

func ExampleSprintStackTrace() {
	err := errors.New("test")
	fmt.Println(errors.SprintStackTrace(err))
}

func Example_logger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := errors.New("test")

	logger.Error("an error occurred", "message", err, "stack", errors.SprintStackTrace(err))
}
