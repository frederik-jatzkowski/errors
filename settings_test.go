package errors_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func ExampleGlobalFormatSettings() {
	// set formatting options globally
	errors.GlobalFormatSettings(
		errors.WithAdvancedFormattingOfExternalErrors(),
		errors.WithStrippedFileNamePrefix("github.com/frederik-jatzkowski/errors/"),
		errors.WithStrippedFuncNamePrefix("github.com/frederik-jatzkowski/errors/"),
	)

	fmt.Printf("%s", errors.New("something happened"))
	// Output:
	// something happened
}
