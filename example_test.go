package errors_test

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/frederik-jatzkowski/errors"
)

func Example_onlyHumanReadable() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Errorf(
			"processing id %d:%w",
			123,
			errors.Join(
				errors.Errorf("two things failed: %w, %w",
					errors.New("something bad happened"),
					errors.Errorf(
						"doing somthing: %w",
						errors.Errorf("failed"),
					),
				),
				errors.New("something else happened"),
			),
		),
	)

	fmt.Printf("%s", err)
}

func Example_includeStackTraces() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Errorf(
			"processing id %d:%w",
			123,
			errors.Join(
				errors.Errorf("two things failed: %w, %w",
					errors.New("something bad happened"),
					errors.Errorf(
						"doing somthing: %w",
						errors.Errorf("failed"),
					),
				),
				errors.New("something else happened"),
			),
		),
	)

	fmt.Printf("%+v", err)
}

func Example_logger() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := errors.New("test")

	logger.Error("an error occurred", "error", err, "full", fmt.Sprintf("%+v", err))
}
