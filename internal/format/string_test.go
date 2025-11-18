package format_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/frederik-jatzkowski/errors/internal/format"
)

func TestString_ProceedToNextError(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		prefix, err, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("").Runes(), nil)

		assert.Equal(t, "", prefix)
		assert.Equal(t, nil, err)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("just text", func(t *testing.T) {
		prefix, err, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc").Runes(), nil)

		assert.Equal(t, "abc", prefix)
		assert.Equal(t, nil, err)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("simple", func(t *testing.T) {
		prefix, err, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %d").Runes(), []any{1})

		assert.Equal(t, "abc 1", prefix)
		assert.Equal(t, nil, err)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("too many args", func(t *testing.T) {
		prefix, err, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc").Runes(), []any{1})

		assert.Equal(t, "abc%!(EXTRA int=1)", prefix)
		assert.Equal(t, nil, err)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("too few args", func(t *testing.T) {
		prefix, err, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %d").Runes(), nil)

		assert.Equal(t, "abc %!d(MISSING)", prefix)
		assert.Equal(t, nil, err)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("only error", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("%w").Runes(), []any{errIn})

		assert.Equal(t, "", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("error at start", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("%w abc").Runes(), []any{errIn})

		assert.Equal(t, "", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Equal(t, format.String(" abc").Runes(), remainingFormat)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("error at end", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %w").Runes(), []any{errIn})

		assert.Equal(t, "abc ", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("not an error", func(t *testing.T) {
		errIn := "not an error"
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %w").Runes(), []any{errIn})

		assert.Equal(t, "abc %!w(string=not an error)", prefix)
		assert.ErrorIs(t, errOut, nil)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("arg before", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("%s abc %w").Runes(), []any{"test", errIn})

		assert.Equal(t, "test abc ", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("args before", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("%s %d abc %w").Runes(), []any{"test", 1, errIn})

		assert.Equal(t, "test 1 abc ", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Len(t, remainingFormat, 0)
		assert.Len(t, remainingArgs, 0)
	})

	t.Run("arg after", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %w %s").Runes(), []any{errIn, "test"})

		assert.Equal(t, "abc ", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Equal(t, format.String(" %s").Runes(), remainingFormat)
		assert.Equal(t, remainingArgs, []any{"test"})
	})

	t.Run("args after", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("abc %w %s %d").Runes(), []any{errIn, "test", 1})

		assert.Equal(t, "abc ", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Equal(t, format.String(" %s %d").Runes(), remainingFormat)
		assert.Equal(t, remainingArgs, []any{"test", 1})
	})

	t.Run("escaped %", func(t *testing.T) {
		errIn := errors.New("hi")
		prefix, errOut, remainingFormat, remainingArgs := format.String("").
			ProceedToNextError(format.String("%% abc %%%w %s %d").Runes(), []any{errIn, "test", 1})

		assert.Equal(t, "% abc %", prefix)
		assert.ErrorIs(t, errOut, errIn)
		assert.Equal(t, format.String(" %s %d").Runes(), remainingFormat)
		assert.Equal(t, remainingArgs, []any{"test", 1})
	})
}

func TestString_SplitIntoComponents(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		err := errors.New("hi")
		components := format.String("%w").SplitIntoComponents([]any{err})
		require.Len(t, components.Components, 1)
		require.Len(t, components.Errs, 1)
	})

	t.Run("empty", func(t *testing.T) {
		components := format.String("").SplitIntoComponents(nil)
		require.Len(t, components.Components, 1)
		require.Len(t, components.Errs, 0)
	})

	t.Run("single wrap", func(t *testing.T) {
		err := errors.New("hi")
		components := format.String("getting id %d: %w").SplitIntoComponents([]any{123, err})
		require.Len(t, components.Components, 2)
		require.Len(t, components.Errs, 1)
	})
}
