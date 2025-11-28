package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
	"github.com/frederik-jatzkowski/errors/examples/nested/subpackage"
)

func main() {
	errors.GlobalFormatSettings(
		errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
		errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
	)

	err := errors.Errorf(
		"call failed: %w",
		errors.Errorf(
			"processing id %d: %w",
			123,
			errors.Join(
				errors.Errorf("double errorf: %w, %w",
					subpackage.SomethingBad(),
					errors.Errorf(
						"hi, %w",
						errors.Errorf("abc"),
					),
				),
				subpackage.SomethingElse(),
			),
		),
	)

	fmt.Printf("%v", err)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%+v", err)
}
