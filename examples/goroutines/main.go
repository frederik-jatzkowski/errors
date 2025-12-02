package main

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func main() {
	errors.GlobalFormatSettings(
		errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
		errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
	)

	err := <-FromGoroutine()

	fmt.Printf("%+v", errors.Errorf("returned from channel: %w", err)) // no additional stack
	fmt.Println()
	fmt.Println()
	fmt.Printf(
		"%+v",
		errors.Errorf("returned from channel: %w", errors.WithStack(err)),
	) // stack before wrapping
	fmt.Println()
	fmt.Println()
	fmt.Printf(
		"%+v",
		errors.WithStack(errors.Errorf("returned from channel: %w", err)),
	) // wrapping before stack
}

func FromGoroutine() chan error {
	c := make(chan error)
	go func() {
		c <- errors.New("from goroutine")
		close(c)
	}()

	return c
}
