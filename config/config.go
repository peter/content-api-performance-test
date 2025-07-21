package config

import (
	"context"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	once   sync.Once
)

// GetLogger returns a singleton instance of the zap logger
func GetLogger() *zap.Logger {
	once.Do(func() {
		var err error

		// Configure zap to log only to stdout
		config := zap.NewProductionConfig()
		config.OutputPaths = []string{"stdout"}
		config.ErrorOutputPaths = []string{"stdout"}
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

		logger, err = config.Build()
		if err != nil {
			panic("Failed to initialize zap logger: " + err.Error())
		}
	})
	return logger
}

// CloseLogger properly syncs the logger when the application shuts down
func CloseLogger() {
	if logger != nil {
		logger.Sync()
	}
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	ConnectionString string
	MaxConns         int
	MinConns         int
}

func LoadDatabaseConfig() *DatabaseConfig {
	log := GetLogger()

	// Check if DATABASE_URL environment variable is set (highest priority)
	if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
		log.Info("Using DATABASE_URL from environment",
			zap.String("database_url", dbURL),
		)
		return &DatabaseConfig{
			ConnectionString: dbURL,
			MaxConns:         getEnvInt("DATABASE_MAX_CONNS", 50),
			MinConns:         getEnvInt("DATABASE_MIN_CONNS", 5),
		}
	}

	engine := GetDatabaseEngine()
	switch engine {
	case "postgres":
		// Sensible default PostgreSQL URL
		defaultPostgresURL := "postgres://postgres:postgres@localhost:5432/content_api?sslmode=disable"
		log.Info("Using default PostgreSQL URL",
			zap.String("database_url", defaultPostgresURL),
		)
		return &DatabaseConfig{
			ConnectionString: defaultPostgresURL,
			MaxConns:         getEnvInt("DATABASE_MAX_CONNS", 50),
			MinConns:         getEnvInt("DATABASE_MIN_CONNS", 5),
		}
	case "sqlite":
		fallthrough
	default:
		// Default SQLite path
		defaultDBPath := "db/sqlite/content-api.db?_journal=WAL"

		// Use default path and ensure directory exists
		if err := os.MkdirAll(filepath.Dir(defaultDBPath), 0755); err != nil {
			log.Warn("Could not create database directory",
				zap.Error(err),
				zap.String("path", filepath.Dir(defaultDBPath)),
			)
		}
		log.Info("Using default SQLite database path",
			zap.String("database_path", defaultDBPath),
		)

		return &DatabaseConfig{
			ConnectionString: defaultDBPath,
			MaxConns:         getEnvInt("DATABASE_MAX_CONNS", 50),
			MinConns:         getEnvInt("DATABASE_MIN_CONNS", 5),
		}
	}
}

// getEnvInt gets an environment variable as an integer with a default value
func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// GetConnectionString returns the database connection string
func (c *DatabaseConfig) GetConnectionString() string {
	return c.ConnectionString
}

// GetDatabaseEngine returns the database engine ("sqlite" or "postgres") based on the DATABASE_ENGINE env var.
func GetDatabaseEngine() string {
	log := GetLogger()
	engine := os.Getenv("DATABASE_ENGINE")
	switch engine {
	case "postgres":
		return "postgres"
	case "sqlite", "":
		return "sqlite"
	default:
		log.Warn("Unknown DATABASE_ENGINE, defaulting to sqlite",
			zap.String("engine", engine),
		)
		return "sqlite"
	}
}

// GetLoggerWithRequestID returns a logger with request ID from context if available
func GetLoggerWithRequestID(ctx context.Context) *zap.Logger {
	logger := GetLogger()

	// Try to get request ID from context
	if requestID := GetRequestIDFromContext(ctx); requestID != "" {
		return logger.With(zap.String("request_id", requestID))
	}

	return logger
}

// GetRequestIDFromContext extracts request ID from context
func GetRequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	// Try to get request ID from context using the middleware RequestIDKey
	if requestID, ok := ctx.Value("request_id").(string); ok {
		return requestID
	}

	return ""
}
