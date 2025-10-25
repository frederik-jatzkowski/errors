package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// testCauseError implements the Cause() interface for testing
type testCauseError struct {
	cause error
	msg   string
}

func (e *testCauseError) Error() string {
	return e.msg
}

func (e *testCauseError) Cause() error {
	return e.cause
}

func TestCause(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := Cause(nil)
		assert.Nil(t, result)
	})

	t.Run("error without cause", func(t *testing.T) {
		err := New("test error")
		result := Cause(err)
		assert.Equal(t, err, result)
	})

	t.Run("error with cause interface", func(t *testing.T) {
		innerErr := New("inner error")
		outerErr := &testCauseError{msg: "outer error", cause: innerErr}

		result := Cause(outerErr)
		assert.Equal(t, innerErr, result)
	})
}

func TestWithMessage(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := WithMessage(nil, "message")
		assert.Nil(t, result)
	})

	t.Run("wrap error with message", func(t *testing.T) {
		originalErr := New("original error")
		wrappedErr := WithMessage(originalErr, "additional context")

		assert.NotNil(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "additional context")
		assert.Contains(t, wrappedErr.Error(), "original error")

		// Should be able to unwrap to get original error
		assert.True(t, Is(wrappedErr, originalErr))
	})
}

func TestWithMessagef(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := WithMessagef(nil, "formatted %s", "message")
		assert.Nil(t, result)
	})

	t.Run("wrap error with formatted message", func(t *testing.T) {
		originalErr := New("original error")
		wrappedErr := WithMessagef(originalErr, "context %d: %s", 42, "additional info")

		assert.NotNil(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "context 42: additional info")
		assert.Contains(t, wrappedErr.Error(), "original error")

		// Should be able to unwrap to get original error
		assert.True(t, Is(wrappedErr, originalErr))
	})
}

func TestWithStack(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := WithStack(nil)
		assert.Nil(t, result)
	})

	t.Run("add stack to error", func(t *testing.T) {
		originalErr := fmt.Errorf("standard error") // nolint: forbidigo
		stackErr := WithStack(originalErr)

		assert.NotNil(t, stackErr)
		// The WithStack implementation uses Join which adds formatting, so we check the content
		assert.Contains(t, stackErr.Error(), originalErr.Error())

		// Should be able to unwrap to get original error
		assert.True(t, Is(stackErr, originalErr))

		// Should have stack trace when formatted with %+v
		stackStr := fmt.Sprintf("%+v", stackErr)
		assert.NotEmpty(t, stackStr)
	})
}

func TestWrap(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := Wrap(nil, "message")
		assert.Nil(t, result)
	})

	t.Run("wrap error with message", func(t *testing.T) {
		originalErr := fmt.Errorf("original error") // nolint: forbidigo
		wrappedErr := Wrap(originalErr, "wrapper message")

		assert.NotNil(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "wrapper message")
		assert.Contains(t, wrappedErr.Error(), "original error")

		// Should be able to unwrap to get original error
		assert.True(t, Is(wrappedErr, originalErr))
	})
}

func TestWrapf(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := Wrapf(nil, "formatted %s", "message")
		assert.Nil(t, result)
	})

	t.Run("wrap error with formatted message", func(t *testing.T) {
		originalErr := fmt.Errorf("original error") // nolint: forbidigo
		wrappedErr := Wrapf(originalErr, "wrapper %d: %s", 123, "context")

		assert.NotNil(t, wrappedErr)
		assert.Contains(t, wrappedErr.Error(), "wrapper 123: context")
		assert.Contains(t, wrappedErr.Error(), "original error")

		// Should be able to unwrap to get original error
		assert.True(t, Is(wrappedErr, originalErr))
	})
}
