package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	config "github.com/ingka-group-digital/b2b-service-pmp/configs"
	"golang.org/x/exp/slog"
)

// global middleware
func ContextMW(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userAgent := r.Header.Get("user-agent")
		transactionId := r.Header.Get("x-transaction-id")

		if transactionId == "" {
			transactionId = uuid.New().String()
		}

		logger := slog.With("transactionId", transactionId, "method", r.Method, "url", r.URL.Path)

		headers := make(map[string]string)
		for key, values := range r.Header {
			headers[key] = values[0]
		}

		message := fmt.Sprintf("➜[%s] %s user-agent %s", r.Method, r.URL.Path, userAgent)
		logger.Info(message,
			"remoteAddr", r.RemoteAddr,
			"headers", headers,
		)

		w.Header().Set("X-Transaction-Id", transactionId)

		ctx := context.WithValue(r.Context(), "transactionId", transactionId)
		ctx = context.WithValue(ctx, "logger", logger)

		//set timeout for the request
		ctx, cancel := context.WithTimeout(ctx, time.Second*config.Get().Timeout)
		defer cancel()

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
