package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors"
)

func TestJoin(t *testing.T) {
	var (
		Err1 = errors.Errorf("error 1")
		Err2 = errors.New("error 2")
	)

	err := errors.Join(Err1, Err2)
	result := fmt.Sprintf("%+v", err)
	assert.Contains(t, result, "error 1")
	assert.Contains(t, result, "error 2")
	assert.ErrorIs(t, err, Err1)
	assert.ErrorIs(t, err, Err2)
	assert.ErrorIs(t, err, err)
	assert.False(t, errors.Is(err, ErrSentinel))
}

func TestJoinError(t *testing.T) {
	err1 := errors.New("first")
	err2 := errors.New("second")
	joined := errors.Join(err1, err2)

	errorMsg := joined.Error()
	assert.NotEmpty(t, errorMsg)
}

func TestJoinNilErrors(t *testing.T) {
	err1 := errors.New("valid error")
	joined := errors.Join(nil, err1, nil)

	assert.ErrorIs(t, joined, err1)
	assert.NotNil(t, joined)
}

func TestJoinAllNil(t *testing.T) {
	result := errors.Join(nil, nil, nil)
	assert.Nil(t, result)
}

func TestJoinSingle(t *testing.T) {
	err := errors.New("single")
	joined := errors.Join(err)
	assert.ErrorIs(t, joined, err)
}

func TestJoinUnwrap(t *testing.T) {
	err1 := errors.New("first")
	err2 := errors.New("second")
	joined := errors.Join(err1, err2)

	if unwrapper, ok := joined.(interface{ Unwrap() []error }); ok {
		unwrapped := unwrapper.Unwrap()
		assert.Len(t, unwrapped, 2)
		assert.Contains(t, unwrapped, err1)
		assert.Contains(t, unwrapped, err2)
	}
}
