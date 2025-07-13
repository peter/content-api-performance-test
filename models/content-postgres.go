package models

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresContentStore struct {
	pool *pgxpool.Pool
}

func NewPostgresContentStore(connString string) (*PostgresContentStore, error) {
	return NewPostgresContentStoreWithConfig(connString, 50, 5)
}

func NewPostgresContentStoreWithConfig(connString string, maxConns, minConns int) (*PostgresContentStore, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("failed to parse connection string: %w", err)
	}

	// Configure pool for high concurrency
	config.MaxConns = int32(maxConns)         // Maximum number of connections in pool
	config.MinConns = int32(minConns)         // Minimum number of connections to maintain
	config.MaxConnLifetime = time.Hour        // Maximum lifetime of a connection
	config.MaxConnIdleTime = 30 * time.Minute // Maximum idle time of a connection
	config.HealthCheckPeriod = time.Minute    // How often to check connection health

	// Create the pool with configured settings
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	// Test the connection
	if err := pool.Ping(context.Background()); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &PostgresContentStore{pool: pool}, nil
}

// Close closes the database connection pool
func (cs *PostgresContentStore) Close() error {
	cs.pool.Close()
	return nil
}

// Create inserts a new content record
func (cs *PostgresContentStore) Create(content *Content) error {
	// Serialize data to JSON
	dataJSON, err := json.Marshal(content.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	now := time.Now()
	content.CreatedAt = now
	content.UpdatedAt = now

	query := `
		INSERT INTO content (id, title, body, author, status, data, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err = cs.pool.Exec(context.Background(), query,
		content.ID,
		content.Title,
		content.Body,
		content.Author,
		content.Status,
		dataJSON,
		content.CreatedAt,
		content.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("failed to create content: %w", err)
	}

	return nil
}

// GetByID retrieves content by ID
func (cs *PostgresContentStore) GetByID(id string) (*Content, error) {
	query := `
		SELECT id, title, body, author, status, data, created_at, updated_at
		FROM content WHERE id = $1
	`

	var content Content
	var dataJSON []byte

	err := cs.pool.QueryRow(context.Background(), query, id).Scan(
		&content.ID,
		&content.Title,
		&content.Body,
		&content.Author,
		&content.Status,
		&dataJSON,
		&content.CreatedAt,
		&content.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("content not found")
		}
		return nil, fmt.Errorf("failed to get content: %w", err)
	}

	// Deserialize data from JSON
	if err := json.Unmarshal(dataJSON, &content.Data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return &content, nil
}

// List retrieves all content records
func (cs *PostgresContentStore) List() ([]*Content, error) {
	query := `
		SELECT id, title, body, author, status, data, created_at, updated_at
		FROM content ORDER BY created_at DESC LIMIT 100
	`

	rows, err := cs.pool.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to query content: %w", err)
	}
	defer rows.Close()

	var contents []*Content
	for rows.Next() {
		var content Content
		var dataJSON []byte

		err := rows.Scan(
			&content.ID,
			&content.Title,
			&content.Body,
			&content.Author,
			&content.Status,
			&dataJSON,
			&content.CreatedAt,
			&content.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan content: %w", err)
		}

		// Deserialize data from JSON
		if err := json.Unmarshal(dataJSON, &content.Data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal data: %w", err)
		}

		contents = append(contents, &content)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return contents, nil
}

// Update updates an existing content record
func (cs *PostgresContentStore) Update(content *Content) error {
	// Serialize data to JSON
	dataJSON, err := json.Marshal(content.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	content.UpdatedAt = time.Now()

	query := `
		UPDATE content 
		SET title = $1, body = $2, author = $3, status = $4, data = $5, updated_at = $6
		WHERE id = $7
	`

	result, err := cs.pool.Exec(context.Background(), query,
		content.Title,
		content.Body,
		content.Author,
		content.Status,
		dataJSON,
		content.UpdatedAt,
		content.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update content: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("content not found")
	}

	return nil
}

// Delete removes a content record by ID
func (cs *PostgresContentStore) Delete(id string) error {
	query := `DELETE FROM content WHERE id = $1`

	result, err := cs.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete content: %w", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("content not found")
	}

	return nil
}
