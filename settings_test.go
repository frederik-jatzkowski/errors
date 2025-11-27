package errors_test

import (
	"fmt"

	"github.com/frederik-jatzkowski/errors"
)

func ExampleGlobalFormatSettings() {
	errors.GlobalFormatSettings(
		errors.WithAdvancedFormattingOfExternalErrors(),
	)

	fmt.Println(errors.New("something happened"))
	// Output:
	// something happened
}
