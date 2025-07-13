package middleware

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/seenthis-ab/content-api/config"
	"go.uber.org/zap"
)

// RequestIDKey is the context key for storing request ID
const RequestIDKey = "request_id"

// GetRequestID retrieves the request ID from the context
func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}

// WithRequestID adds a request ID to the context
func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

// ResponseTimeWriter wraps http.ResponseWriter to add X-Response-Time header
type ResponseTimeWriter struct {
	http.ResponseWriter
	startTime  time.Time
	statusCode int
	logger     *zap.Logger
	requestID  string
}

// WriteHeader adds the X-Response-Time header before writing the status code
func (w *ResponseTimeWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	duration := time.Since(w.startTime)
	responseTime := float64(duration.Microseconds()) / 1000.0
	w.Header().Set("X-Response-Time", strconv.FormatFloat(responseTime, 'f', 2, 64)+"ms")

	w.logger.Debug("Setting response time header",
		zap.Float64("response_time_ms", responseTime),
		zap.String("request_id", w.requestID),
	)
	w.ResponseWriter.WriteHeader(statusCode)
}

// Write delegates to the underlying ResponseWriter
func (w *ResponseTimeWriter) Write(data []byte) (int, error) {
	return w.ResponseWriter.Write(data)
}

// LoggingMiddleware creates a middleware that logs HTTP requests and responses
func LoggingMiddleware() func(http.Handler) http.Handler {
	logger := config.GetLogger()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Generate a lowercase ULID for request ID
			requestID := strings.ToLower(ulid.Make().String())

			// Add requestID to the request context
			ctx := WithRequestID(r.Context(), requestID)
			r = r.WithContext(ctx)

			// Log request details
			logger.Info("HTTP request started",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("user_agent", r.UserAgent()),
				zap.String("remote_addr", r.RemoteAddr),
				zap.String("request_id", requestID),
				zap.Time("timestamp", start),
			)

			// Create a custom response writer to capture headers and status
			responseWriter := &ResponseTimeWriter{
				ResponseWriter: w,
				startTime:      start,
				logger:         logger,
				requestID:      requestID,
			}

			// Set the X-Request-Id header
			responseWriter.Header().Set("X-Request-Id", requestID)

			next.ServeHTTP(responseWriter, r)

			duration := time.Since(start)
			responseTime := float64(duration.Microseconds()) / 1000.0

			// Log response details
			logger.Info("HTTP request completed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Int("status_code", responseWriter.statusCode),
				zap.Float64("duration_ms", responseTime),
				zap.String("content_length", responseWriter.Header().Get("Content-Length")),
				zap.String("request_id", requestID),
				zap.Time("timestamp", time.Now()),
			)
		})
	}
}
