package models

import (
	"fmt"
	"time"

	"github.com/seenthis-ab/content-api/config"
	"go.uber.org/zap"
)

// Content represents a content record
type Content struct {
	ID        string                 `json:"id" db:"id"`
	Title     string                 `json:"title" db:"title"`
	Body      string                 `json:"body" db:"body"`
	Author    string                 `json:"author" db:"author"`
	Status    string                 `json:"status" db:"status"`
	Data      map[string]interface{} `json:"data" db:"data"`
	CreatedAt time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt time.Time              `json:"updated_at" db:"updated_at"`
}

// ContentStore interface defines the contract for content storage operations
type ContentStore interface {
	Create(content *Content) error
	GetByID(id string) (*Content, error)
	List() ([]*Content, error)
	Update(content *Content) error
	Delete(id string) error
	Close() error
}

// ContentStoreFactory defines a function type for creating ContentStore instances
type ContentStoreFactory func(dbConfig *config.DatabaseConfig) (ContentStore, error)

// GetContentStore creates and returns the appropriate ContentStore implementation
// based on the database configuration
func GetContentStore() (ContentStore, error) {
	log := config.GetLogger()
	dbConfig := config.LoadDatabaseConfig()
	switch config.GetDatabaseEngine() {
	case "postgres":
		log.Info("Using Postgres database engine",
			zap.Int("max_connections", dbConfig.MaxConns),
			zap.Int("min_connections", dbConfig.MinConns),
		)
		pgStore, err := NewPostgresContentStoreWithConfig(dbConfig.GetConnectionString(), dbConfig.MaxConns, dbConfig.MinConns)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize Postgres database: %w", err)
		}
		return pgStore, nil
	case "sqlite":
		log.Info("Using SQLite database engine")
		fallthrough
	default:
		sqlStore, err := NewSQLiteContentStore(dbConfig.GetConnectionString())
		if err != nil {
			return nil, fmt.Errorf("failed to initialize SQLite database: %w", err)
		}
		return sqlStore, nil
	}
}
