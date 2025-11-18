package dto

import (
	"io"
	"strings"
)

type Writer struct {
	inner          io.Writer
	isAfterNewline bool
	depth          int
}

func NewWriter(inner io.Writer, initialDepth int) *Writer {
	return &Writer{inner: inner, depth: initialDepth}
}

var _ io.Writer = &Writer{}

func (w *Writer) Write(p []byte) (n int, err error) {
	lines := strings.Split(string(p), "\n")
	for i, line := range lines {
		if i < len(lines)-1 {
			w.AddNewline()
		}
		w.Indent()

		_, err := w.inner.Write([]byte(line))
		if err != nil {
			return 0, err
		}

		w.isAfterNewline = true
	}

	w.isAfterNewline = strings.HasSuffix(string(p), "\n")

	return len(p), nil
}

func (w *Writer) IsAfterNewline() bool {
	return w.isAfterNewline
}

func (w *Writer) AddNewline() {
	// nolint: errcheck
	w.inner.Write([]byte{'\n'})
	w.isAfterNewline = true
}

func (w *Writer) Descend() {
	w.depth++
}

func (w *Writer) Ascend() {
	w.depth--
}

func (w *Writer) Indent() {
	if w.isAfterNewline {
		// nolint: errcheck
		w.inner.Write([]byte(strings.Repeat("    ", max(0, w.depth))))
	}
}
