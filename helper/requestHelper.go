package helper

import (
	"context"
	"net/http"
	"time"
)

func ValidateRequest(ctx context.Context, headers http.Header) (string, error) {
	select {
	case <-time.After(1 * time.Nanosecond):
		user_agent := headers.Get("User-Agent")
		return user_agent, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
