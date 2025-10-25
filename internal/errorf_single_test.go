package internal_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors/internal"
)

func Test_errorfSingle(t *testing.T) {
	self := &internal.ErrorfSingle{}
	with := &internal.WithStack{}

	t.Run("As", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = internal.EnsureStackTraceIfNecessary(2, &internal.ErrorfSingle{
			Msg:     "hi",
			Wrapped: err,
		}, []error{err})

		assert.ErrorAs(t, err, &self)
		assert.ErrorAs(t, err, &with)
		assert.NotEqual(t, err.Error(), fmt.Sprintf("%+v", err))
	})

	t.Run("Is", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = internal.EnsureStackTraceIfNecessary(2, &internal.ErrorfSingle{
			Msg:     "hi",
			Wrapped: err,
		}, []error{err})

		assert.ErrorIs(t, err, err)
	})

	t.Run("As edge cases", func(t *testing.T) {
		single := &internal.ErrorfSingle{
			Msg:     "test error",
			Wrapped: errors.New("inner"), // nolint: forbidigo
		}

		// Test As with unsupported target type
		var stringPtr *string
		assert.False(t, single.As(&stringPtr))

		// Test As with nil withStack
		single.Stack = nil
		var withStackPtr *internal.WithStack
		assert.False(t, single.As(&withStackPtr))
	})
}
