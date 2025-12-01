package main

import (
	"log/slog"
	"os"

	"github.com/frederik-jatzkowski/errors"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	err := errors.New("test")

	// All errors from this package implement slog.LogValuer.
	// They will format itself into an object shaped like {"short": "...", "long": "... with stacks"}.
	logger.Error("an error occurred", "error", err)
}
