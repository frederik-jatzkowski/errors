package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

var (
	ErrSomethingWentWrong    = errors.New("something went wrong")
	ErrSomthingElseWentWrong = errors.New("something else went wrong")
)

func main() {
	err := errors.Errorf(
		"call failed: %w",
		errors.Join(
			ErrSomethingWentWrong,
			errors.New("something bad happened"),
			ErrSomthingElseWentWrong,
		),
	)

	fmt.Printf("%+v\n", err)
	fmt.Printf("%+v\n", errors.Errorf("abc"))
}
