package logger

import (
	"os"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	"golang.org/x/exp/slog"
)

var logLevelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
	"trace": slog.LevelDebug,
	"fatal": slog.LevelError,
}

func parseLogLevel(level string) slog.Level {
	if logLevel, ok := logLevelMap[level]; ok {
		return logLevel
	}
	return slog.LevelInfo //default
}

func Init(namespace *string) {
	logLevel := parseLogLevel(config.Get().LogLevel)

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	})

	logger := slog.New(handler)

	if namespace != nil {
		logger = logger.With("ns", *namespace)
	}

	slog.SetDefault(logger)
}
