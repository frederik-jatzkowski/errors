// Package internal holds the internal implementations for [github.com/frederik-jatzkowski/errors]
package internal

type Error interface {
	error
	SetWithStack(ws *WithStack)
}
