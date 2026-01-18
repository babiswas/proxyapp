package helper

import (
	"app/logger"
	"context"
	"net/http"
	"time"
)

func ValidateRequest(ctx context.Context, headers http.Header) (string, error) {
	logger := logger.LoggingInit()
	logger.Info("Executing request validation block.")
	select {
	case <-time.After(1 * time.Nanosecond):
		logger.Info("Fetching request headers to identify user agent.")
		user_agent := headers.Get("User-Agent")
		return user_agent, nil
	case <-ctx.Done():
		logger.Info("Request deadline exceeded.")
		return "", ctx.Err()
	}
}
