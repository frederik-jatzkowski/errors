package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors"
)

var ErrSentinel = errors.New("sentinel error")

func SomeFunction12345() error {
	return errors.New("test")
}

func TestNew(t *testing.T) {
	result := fmt.Sprintf("%+v", ErrSentinel)
	assert.Equal(t, "sentinel error", result)

	result = fmt.Sprintf("%+v", SomeFunction12345())
	assert.Contains(t, result, "test")
	assert.Contains(t, result, "github.com/frederik-jatzkowski/errors_test.TestNew")

	result = errors.Errorf("%+v", errors.New("hi")).Error()
	assert.Contains(t, result, "hi")
	assert.Contains(t, result, "github.com/frederik-jatzkowski/errors_test.TestNew")
}

func TestWithStackError(t *testing.T) {
	err := errors.New("base error")
	assert.Equal(t, "base error", err.Error())
}

func TestStackTraceFormat(t *testing.T) {
	err := errors.New("format test")

	result1 := fmt.Sprintf("%v", err)
	assert.Contains(t, result1, "format test")

	result2 := fmt.Sprintf("%s", err)
	assert.Contains(t, result2, "format test")

	result3 := fmt.Sprintf("%q", err)
	assert.Contains(t, result3, "format test")
}

func TestWrapAs(t *testing.T) {
	var target error
	err := errors.New("test")

	result := errors.As(err, &target)
	assert.True(t, result)
	assert.Equal(t, err, target)
}

func TestWrapUnwrap(t *testing.T) {
	base := errors.New("base")
	wrapped := errors.Errorf("wrapped: %w", base)

	result := errors.Unwrap(wrapped)
	assert.Equal(t, base, result)

	assert.NotNil(t, result)
}

func TestErrorFormatting(t *testing.T) {
	err := errors.New("formatting test")

	result1 := fmt.Sprintf("%+v", err)
	assert.Contains(t, result1, "formatting test")

	result2 := fmt.Sprintf("%-v", err)
	assert.Contains(t, result2, "formatting test")

	result3 := fmt.Sprintf("%#v", err)
	assert.Contains(t, result3, "formatting test")
}

func TestErrorIsAndAs(t *testing.T) {
	err1 := errors.New("original")
	wrapped := errors.Errorf("wrapped: %w", err1)

	assert.True(t, errors.Is(wrapped, err1))
	assert.False(t, errors.Is(wrapped, errors.New("different")))

	var target error
	assert.True(t, errors.As(wrapped, &target))
}

func TestComplexErrorChain(t *testing.T) {
	base := errors.New("base error")
	wrapped1 := errors.Errorf("level 1: %w", base)
	wrapped2 := errors.Errorf("level 2: %w", wrapped1)
	joined := errors.Join(wrapped2, errors.New("additional"))

	assert.ErrorIs(t, joined, base)
	assert.Contains(t, joined.Error(), "level 2")
}

func TestDirectErrorMethods(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	manyErr := errors.Errorf("test %w and %w", err1, err2)

	result := manyErr.Error()
	assert.Contains(t, result, "test")

	singleErr := errors.Errorf("single %w", err1)
	result2 := singleErr.Error()
	assert.Contains(t, result2, "single")

	joined := errors.Join(err1, err2)
	result3 := joined.Error()
	assert.NotEmpty(t, result3)
}

func TestStackTraceString(t *testing.T) {
	err := errors.New("string test")
	result := fmt.Sprintf("%s", err)
	assert.Contains(t, result, "string test")
}

func TestSpecificFormatting(t *testing.T) {
	err1 := errors.New("format1")
	err2 := errors.New("format2")
	joined := errors.Join(err1, err2)

	result1 := fmt.Sprintf("%+v", joined)
	assert.NotEmpty(t, result1)

	result2 := fmt.Sprintf("%#v", joined)
	assert.NotEmpty(t, result2)

	result3 := fmt.Sprintf("%T", joined)
	assert.NotEmpty(t, result3)
}

func TestErrorfManyFormat(t *testing.T) {
	err1 := errors.New("first")
	err2 := errors.New("second")
	manyErr := errors.Errorf("many: %w and %w", err1, err2)

	result := fmt.Sprintf("%+v", manyErr)
	assert.Contains(t, result, "many:")
	assert.NotEmpty(t, result)
}

func TestStackTraceStringMethod(t *testing.T) {
	err := errors.New("string method test")

	result := fmt.Sprintf("%s", err)
	assert.Contains(t, result, "string method test")
}

func TestStackTraceFormatEdgeCases(t *testing.T) {
	err := errors.New("format edge test")

	result1 := fmt.Sprintf("%v", err) // basic verb
	assert.NotEmpty(t, result1)

	result2 := fmt.Sprintf("%.10v", err) // with precision
	assert.NotEmpty(t, result2)

	result3 := fmt.Sprintf("%10v", err) // with width
	assert.NotEmpty(t, result3)

	result4 := fmt.Sprintf("%-10v", err) // with left align flag
	assert.NotEmpty(t, result4)
}

func TestFormatFrameEdgeCases(t *testing.T) {
	err := errors.New("frame test")

	result1 := fmt.Sprintf("%+s", err) // string with plus flag
	assert.NotEmpty(t, result1)

	result2 := fmt.Sprintf("%#s", err) // string with hash flag
	assert.NotEmpty(t, result2)

	result3 := fmt.Sprintf("%+d", err) // different verb with plus
	assert.NotEmpty(t, result3)
}

func TestStackTraceHashFormat(t *testing.T) {
	err := errors.New("hash format test")

	result := fmt.Sprintf("%#v", err)
	assert.NotEmpty(t, result)
}

func TestDirectStringCall(t *testing.T) {
	err := errors.New("direct string test")

	if stringer, ok := err.(fmt.Stringer); ok {
		result := stringer.String()
		assert.NotEmpty(t, result)
	}
}

func TestErrorfVerbIndexEdge(t *testing.T) {
	err1 := errors.Errorf("no args")
	assert.Equal(t, "no args", err1.Error())

	err2 := errors.Errorf("mixed verbs %s %w %d", "string", errors.New("wrapped"), 42)
	assert.NotEmpty(t, err2.Error())

	err3 := errors.Errorf("%w%w", errors.New("first"), errors.New("second"))
	assert.NotEmpty(t, err3.Error())
}
