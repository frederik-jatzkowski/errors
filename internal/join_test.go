package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors"
	"github.com/frederik-jatzkowski/errors/internal"
)

func Test_join_As_edge_cases(t *testing.T) {
	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	j := &internal.Join{
		Wrapped: []error{err1, err2},
	}

	t.Run("As with join target", func(t *testing.T) {
		var joinPtr *internal.Join
		assert.True(t, j.As(&joinPtr))
		assert.Equal(t, j, joinPtr)
	})

	t.Run("As with withStack target - nil stack", func(t *testing.T) {
		j.Stack = nil
		var withStackPtr *internal.WithStack
		assert.False(t, j.As(&withStackPtr))
	})

	t.Run("As with withStack target - has stack", func(t *testing.T) {
		ws := &internal.WithStack{Inner: j, St: internal.NewStackTrace(4)}
		j.Stack = ws
		var withStackPtr *internal.WithStack
		assert.True(t, j.As(&withStackPtr))
		assert.Equal(t, ws, withStackPtr)
	})

	t.Run("As with unsupported target", func(t *testing.T) {
		var stringPtr *string
		assert.False(t, j.As(&stringPtr))
	})
}
