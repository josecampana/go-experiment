package logger

import (
	"os"
	"sync"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	"golang.org/x/exp/slog"
)

// singleton management
var (
	instance *slog.Logger
	once     sync.Once
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

func Init(namespace *string) *slog.Logger {
	once.Do(func() {
		logLevel := parseLogLevel(config.Get().LogLevel)

		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		})

		logger := slog.New(handler)

		if namespace != nil {
			logger = logger.With("ns", *namespace)
		}

		slog.SetDefault(logger)
		instance = logger
	})

	return instance
}

func Get() *slog.Logger {
	if instance == nil {
		return Init(nil)
	}

	return instance
}

func WithModule(module string) *slog.Logger {
	if instance == nil {
		return Init(&module)
	}

	return instance.With("module", &module)
}


