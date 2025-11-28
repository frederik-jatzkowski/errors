// Package subpackage is a package for demonstrating errors from imported packages
package subpackage

import "github.com/frederik-jatzkowski/errors"

func SomethingBad() error {
	return errors.New("something bad happened")
}

func SomethingElse() error {
	return errors.New("something else happened")
}
