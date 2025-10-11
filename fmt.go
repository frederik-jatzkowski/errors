package errors

import "fmt"

func shouldPrintStack(s fmt.State, verb rune) bool {
	return verb == 'v' && s.Flag('+')
}

func printErrorString(msg string, s fmt.State, verb rune) {
	// nolint: errcheck
	fmt.Fprintf(s, fmt.FormatString(s, verb), msg)
}

func delegateFormat(err error, s fmt.State, verb rune) {
	// nolint: errcheck
	fmt.Fprintf(s, fmt.FormatString(s, verb), err)
}
