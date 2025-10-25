package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors/internal"
)

func Test_simple_setWithStack(t *testing.T) {
	s := &internal.Simple{Msg: "test error"}
	ws := &internal.WithStack{
		Inner: s,
		St:    internal.NewStackTrace(3),
	}

	s.SetWithStack(ws)
	assert.Equal(t, ws, s.Stack)
}

func Test_simple_As_edge_cases(t *testing.T) {
	s := &internal.Simple{Msg: "test error"}

	t.Run("As with simple target", func(t *testing.T) {
		var simplePtr *internal.Simple
		assert.True(t, s.As(&simplePtr))
		assert.Equal(t, s, simplePtr)
	})

	t.Run("As with withStack target - nil stack", func(t *testing.T) {
		s.Stack = nil
		var withStackPtr *internal.WithStack
		assert.False(t, s.As(&withStackPtr))
	})

	t.Run("As with withStack target - has stack", func(t *testing.T) {
		ws := &internal.WithStack{Inner: s, St: internal.NewStackTrace(3)}
		s.Stack = ws
		var withStackPtr *internal.WithStack
		assert.True(t, s.As(&withStackPtr))
		assert.Equal(t, ws, withStackPtr)
	})

	t.Run("As with unsupported target", func(t *testing.T) {
		var stringPtr *string
		assert.False(t, s.As(&stringPtr))
	})
}
