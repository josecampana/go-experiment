package ContextHelper

import (
	"context"

	"golang.org/x/exp/slog"
)

// Define claves de contexto para el transactionId y el logger
const (
	transactionIDKey string = "transactionId"
	loggerKey        string = "logger"
)

func TransactionId(ctx context.Context) string {
	if transactionId, ok := ctx.Value(transactionIDKey).(string); ok {
		return transactionId
	}
	return ""
}

func Logger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
