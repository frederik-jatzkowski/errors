package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	ErrSentinel     = New("sentinel")
	ErrSentinelInit error
)

func init() {
	ErrSentinelInit = New("sentinel2")
}

func Test_stackTrace_isSentinel(t *testing.T) {
	_, ok := ErrSentinel.(*simple)
	require.True(t, ok)
}

func Test_stackTrace_isSentinel_init(t *testing.T) {
	with, ok := ErrSentinelInit.(*withStack)
	require.True(t, ok)
	assert.False(t, with.st.isSentinel())
}

func Test_stackTrace_isSentinel_false(t *testing.T) {
	with, ok := New("test").(*withStack)
	require.True(t, ok)
	assert.False(t, with.st.isSentinel())
}

func Test_stackTrace_String(t *testing.T) {
	with, ok := New("test").(*withStack)
	require.True(t, ok)

	str := with.st.String()
	assert.NotEmpty(t, str)
	assert.Contains(t, str, "Test_stackTrace_String")
}

func Test_stackTrace_isSentinel_edge_cases(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		st := &stackTrace{}
		assert.True(t, st.isSentinel())
	})
}

func Test_stackTrace_Format_edge_cases(t *testing.T) {
	t.Run("format with hash flag", func(t *testing.T) {
		with, ok := New("test").(*withStack)
		require.True(t, ok)

		// Test with # flag to trigger the different format path
		result := fmt.Sprintf("%#v", with.st)
		assert.NotEmpty(t, result)
	})
}

func Test_formatFrame_edge_cases(t *testing.T) {
	t.Run("invalid PC", func(t *testing.T) {
		// Test with invalid PC value (0) to trigger the "unknown" case
		result := formatFrame(0)
		assert.Equal(t, "unknown", result)
	})
}
