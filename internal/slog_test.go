package internal_test

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/frederik-jatzkowski/errors"
)

func TestErrorfMany_LogValue(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info(
		"operation failed",
		"error",
		errors.Errorf("test: %w %w", errors.New("test 1"), errors.New("test 2")),
	)

	assert.Contains(t, buf.String(), "error")
	assert.Contains(t, buf.String(), "short")
	assert.Contains(t, buf.String(), "long")
}

func TestErrorfSingle_LogValue(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info("operation failed", "error", errors.Errorf("test: %w", errors.New("test")))

	assert.Contains(t, buf.String(), "error")
	assert.Contains(t, buf.String(), "short")
	assert.Contains(t, buf.String(), "long")
}

func TestJoin_LogValue(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info("operation failed", "error", errors.Join(errors.New("1"), errors.New("2")))

	assert.Contains(t, buf.String(), "error")
	assert.Contains(t, buf.String(), "short")
	assert.Contains(t, buf.String(), "long")
}

func TestSimple_LogValue(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info("operation failed", "error", errors.New("test"))

	assert.Contains(t, buf.String(), "error")
	assert.Contains(t, buf.String(), "short")
	assert.Contains(t, buf.String(), "long")
}

func TestWithStack_LogValue(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info("operation failed", "error", errors.Errorf("test"))

	assert.Contains(t, buf.String(), "error")
	assert.Contains(t, buf.String(), "short")
	assert.Contains(t, buf.String(), "long")
}
