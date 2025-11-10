package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

var (
	ErrSomethingWentWrong     = errors.New("something went wrong")
	ErrSomethingElseWentWrong = errors.Errorf("something else went wrong")
)

func main() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Join(
			ErrSomethingWentWrong,
			errors.New("something bad happened"),
			ErrSomethingElseWentWrong,
		),
	)

	fmt.Printf("%+v\n", err)
}
