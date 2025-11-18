package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func main() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Errorf(
			"processing id %d: %w",
			123,
			errors.Join(
				errors.Errorf("double errorf: %w, %w",
					errors.New("something bad happened"),
					errors.Errorf(
						"hi, %w",
						errors.Errorf("abc"),
					),
				),
				errors.New("something else happened"),
			),
		),
	)

	fmt.Printf("%v", err)
	fmt.Println()
	fmt.Println()
	fmt.Printf("%+v", err)
}
