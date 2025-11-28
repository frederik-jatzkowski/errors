package subpackage

import "github.com/frederik-jatzkowski/errors"

func SomethingElse() error {
	return errors.New("something else happened")
}
