package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_errorfSingle(t *testing.T) {
	self := &errorfSingle{}
	with := &withStack{}

	t.Run("As", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = ensureStackTraceIfNecessary(&errorfSingle{
			msg:     "hi",
			wrapped: err,
		}, []error{err})

		assert.ErrorAs(t, err, &self)
		assert.ErrorAs(t, err, &with)
		assert.NotEqual(t, err.Error(), fmt.Sprintf("%+v", err))
	})

	t.Run("Is", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = ensureStackTraceIfNecessary(&errorfSingle{
			msg:     "hi",
			wrapped: err,
		}, []error{err})

		assert.ErrorIs(t, err, err)
	})

	t.Run("As edge cases", func(t *testing.T) {
		single := &errorfSingle{
			msg:     "test error",
			wrapped: errors.New("inner"), // nolint: forbidigo
		}

		// Test As with unsupported target type
		var stringPtr *string
		assert.False(t, single.As(&stringPtr))

		// Test As with nil withStack
		single.stack = nil
		var withStackPtr *withStack
		assert.False(t, single.As(&withStackPtr))
	})
}
