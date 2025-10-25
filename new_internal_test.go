package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_simple_setWithStack(t *testing.T) {
	s := &simple{msg: "test error"}
	ws := &withStack{
		inner: s,
		st:    newStackTrace(),
	}

	s.setWithStack(ws)
	assert.Equal(t, ws, s.stack)
}

func Test_simple_As_edge_cases(t *testing.T) {
	s := &simple{msg: "test error"}

	t.Run("As with simple target", func(t *testing.T) {
		var simplePtr *simple
		assert.True(t, s.As(&simplePtr))
		assert.Equal(t, s, simplePtr)
	})

	t.Run("As with withStack target - nil stack", func(t *testing.T) {
		s.stack = nil
		var withStackPtr *withStack
		assert.False(t, s.As(&withStackPtr))
	})

	t.Run("As with withStack target - has stack", func(t *testing.T) {
		ws := &withStack{inner: s, st: newStackTrace()}
		s.stack = ws
		var withStackPtr *withStack
		assert.True(t, s.As(&withStackPtr))
		assert.Equal(t, ws, withStackPtr)
	})

	t.Run("As with unsupported target", func(t *testing.T) {
		var stringPtr *string
		assert.False(t, s.As(&stringPtr))
	})
}
