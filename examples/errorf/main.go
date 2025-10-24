package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func main() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Errorf(
			"doing a: %w, doing b: %w",
			errors.New("something bad happened"),
			fmt.Errorf("external dependency error"), // nolint: forbidigo
		),
	)

	fmt.Printf("%+v", err)
}
