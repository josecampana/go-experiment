package logger

import (
	"os"

	"golang.org/x/exp/slog"
)

// type Fields = slog.Fields

// func NewLogger() *slog.Logger {
// 	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
// }

func Init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	slog.SetDefault(logger)
}
