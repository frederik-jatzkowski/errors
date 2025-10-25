package internal_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors/internal"
)

func Test_errorfMany(t *testing.T) {
	many := &internal.ErrorfMany{}
	with := &internal.WithStack{}

	t.Run("As", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = internal.EnsureStackTraceIfNecessary(2, &internal.ErrorfMany{
			Msg:     "hi",
			Wrapped: []error{err},
		}, []error{err})

		assert.ErrorAs(t, err, &many)
		assert.ErrorAs(t, err, &with)
		assert.NotEqual(t, err.Error(), fmt.Sprintf("%+v", err))
	})

	t.Run("Is", func(t *testing.T) {
		err := errors.New("hello world") // nolint: forbidigo
		err = internal.EnsureStackTraceIfNecessary(2, &internal.ErrorfMany{
			Msg:     "hi",
			Wrapped: []error{err},
		}, []error{err})

		assert.ErrorIs(t, err, err)
	})

	t.Run("Unwrap", func(t *testing.T) {
		err1 := errors.New("error 1") // nolint: forbidigo
		err2 := errors.New("error 2") // nolint: forbidigo
		many := &internal.ErrorfMany{
			Msg:     "multiple errors",
			Wrapped: []error{err1, err2},
		}

		unwrapped := many.Unwrap()
		assert.Len(t, unwrapped, 2)
		assert.Equal(t, err1, unwrapped[0])
		assert.Equal(t, err2, unwrapped[1])
	})

	t.Run("As edge cases", func(t *testing.T) {
		many := &internal.ErrorfMany{
			Msg:     "test error",
			Wrapped: []error{errors.New("inner")}, // nolint: forbidigo
		}

		// Test As with unsupported target type
		var stringPtr *string
		assert.False(t, many.As(&stringPtr))

		// Test As with nil withStack
		many.Stack = nil
		var withStackPtr *internal.WithStack
		assert.False(t, many.As(&withStackPtr))
	})
}
