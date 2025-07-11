package logger

import (
	"log/slog"
	"os"
)

func Set() {
	logger := slog.New(
		slog.NewTextHandler(os.Stderr, nil),
	)

	slog.SetDefault(logger)
}
