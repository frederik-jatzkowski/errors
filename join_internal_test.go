package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_join_As_edge_cases(t *testing.T) {
	err1 := New("error 1")
	err2 := New("error 2")
	j := &join{
		wrapped: []error{err1, err2},
	}

	t.Run("As with join target", func(t *testing.T) {
		var joinPtr *join
		assert.True(t, j.As(&joinPtr))
		assert.Equal(t, j, joinPtr)
	})

	t.Run("As with withStack target - nil stack", func(t *testing.T) {
		j.stack = nil
		var withStackPtr *withStack
		assert.False(t, j.As(&withStackPtr))
	})

	t.Run("As with withStack target - has stack", func(t *testing.T) {
		ws := &withStack{inner: j, st: newStackTrace()}
		j.stack = ws
		var withStackPtr *withStack
		assert.True(t, j.As(&withStackPtr))
		assert.Equal(t, ws, withStackPtr)
	})

	t.Run("As with unsupported target", func(t *testing.T) {
		var stringPtr *string
		assert.False(t, j.As(&stringPtr))
	})
}
