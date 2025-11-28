package errors_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func ExampleGlobalFormatSettings() {
	errors.GlobalFormatSettings(
		errors.WithAdvancedFormattingOfExternalErrors(),
		errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
	)

	fmt.Printf("%+v", errors.New("something happened"))
	// Output:
	// something happened
}
