package internal_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/frederik-jatzkowski/errors"
	"github.com/frederik-jatzkowski/errors/internal"
)

var (
	ErrSentinel     = errors.New("sentinel")
	ErrSentinelInit error
)

func init() {
	ErrSentinelInit = errors.New("sentinel2")
}

func Test_stackTrace_isSentinel(t *testing.T) {
	_, ok := ErrSentinel.(*internal.Simple)
	require.True(t, ok)
}

func Test_stackTrace_isSentinel_init(t *testing.T) {
	with, ok := ErrSentinelInit.(*internal.WithStack)
	require.True(t, ok)
	assert.False(t, with.St.IsSentinel())
}

func Test_stackTrace_isSentinel_false(t *testing.T) {
	with, ok := errors.New("test").(*internal.WithStack)
	require.True(t, ok)
	assert.False(t, with.St.IsSentinel())
}

func Test_stackTrace_isSentinel_edge_cases(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		st := &internal.StackTrace{}
		assert.True(t, st.IsSentinel())
	})
}

func Test_stackTrace_Format_edge_cases(t *testing.T) {
	t.Run("format with hash flag", func(t *testing.T) {
		with, ok := errors.New("test").(*internal.WithStack)
		require.True(t, ok)

		// Test with # flag to trigger the different format path
		result := fmt.Sprintf("%#v", with.St)
		assert.NotEmpty(t, result)
	})
}
