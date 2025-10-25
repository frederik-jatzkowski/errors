package errors

import (
	"errors"
	"fmt"
	"testing"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = New("test")
	}
}

func BenchmarkStdNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = errors.New("test") // nolint: forbidigo
	}
}

func BenchmarkErrorf(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := errors.New("err2") // nolint: forbidigo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Errorf("%s %w %w", "hi", err1, err2)
	}
}

func BenchmarkErrorf_stackExists(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := New("err2")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Errorf("%s %w %w", "hi", err1, err2)
	}
}

func BenchmarkStdErrorf(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := errors.New("err2") // nolint: forbidigo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = fmt.Errorf("%s %w %w", "hi", err1, err2) // nolint: forbidigo
	}
}

func BenchmarkJoin(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := errors.New("err2") // nolint: forbidigo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Join(err1, err2)
	}
}

func BenchmarkJoin_stackExists(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := New("err2")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Join(err1, err2)
	}
}

func BenchmarkStdJoin(b *testing.B) {
	err1 := errors.New("err1") // nolint: forbidigo
	err2 := errors.New("err2") // nolint: forbidigo

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = errors.Join(err1, err2) // nolint: forbidigo
	}
}

func BenchmarkJoin_Deep5(b *testing.B) {
	depth := 5
	err := New("test")
	for i := 0; i < depth; i++ {
		err = Errorf("%d %w", i, err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Join(err)
	}
}

func BenchmarkJoin_Deep50(b *testing.B) {
	depth := 50
	err := New("test")
	for i := 0; i < depth; i++ {
		err = Errorf("%d %w", i, err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Join(err)
	}
}
