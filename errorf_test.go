package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors"
)

func TestErrorf(t *testing.T) {
	errors.Is(errors.New("hi"), errors.New("hi"))
	err := errors.Errorf("some %s: %w", "error", ErrSentinel)
	result := fmt.Sprintf("%+v", err)
	assert.Contains(t, result, "some error: sentinel error")
	assert.Contains(t, result, "github.com/frederik-jatzkowski/errors_test.TestErrorf")
	assert.ErrorIs(t, err, ErrSentinel)
	assert.ErrorIs(t, ErrSentinel, ErrSentinel)
}

func TestErrorfSingle(t *testing.T) {
	err := errors.Errorf("single error: %w", ErrSentinel)
	assert.Equal(t, "single error: sentinel error", err.Error())
	assert.ErrorIs(t, err, ErrSentinel)
}

func TestErrorfMany(t *testing.T) {
	err1 := errors.New("first error")
	err2 := errors.New("second error")
	err := errors.Errorf("many errors: %w and %w", err1, err2)
	assert.Contains(t, err.Error(), "many errors:")
	errorMsg := err.Error()
	assert.NotEmpty(t, errorMsg)
}

func TestErrorfNoWrapping(t *testing.T) {
	err := errors.Errorf("simple error %s", "message")
	assert.Equal(t, "simple error message", err.Error())
}

func TestErrorfMultiplePercent(t *testing.T) {
	err := errors.Errorf("error with %% and %w", ErrSentinel)
	assert.Contains(t, err.Error(), "error with %")
	assert.ErrorIs(t, err, ErrSentinel)
}

func TestErrorfEdgeCases(t *testing.T) {
	err := errors.Errorf("test %w %d %w", ErrSentinel, 42, errors.New("second"))
	assert.Contains(t, err.Error(), "test")
	assert.Contains(t, err.Error(), "42")

	err2 := errors.Errorf("   %w   %s   %w   ", ErrSentinel, "middle", errors.New("end"))
	assert.Contains(t, err2.Error(), "middle")
	assert.NotEmpty(t, err2.Error())
}
