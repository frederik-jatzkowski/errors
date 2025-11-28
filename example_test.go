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
				errors.Errorf("two things failed:%w,%w",
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
	// Output:
	// call failed: processing id 123:
	//     => two things failed:
	//         => something bad happened,
	//         => doing somthing: failed
	//     => something else happened
}

func Example_includeStackTraces_minimal() {
	errors.GlobalFormatSettings(
		errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
	)

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

func Example_slog() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := errors.New("test")

	// All errors from this package implement slog.LogValuer.
	// They will format itself into an object shaped like {"short": "...", "long": "... with stacks"}.
	logger.Error("an error occurred", "error", err)
}
