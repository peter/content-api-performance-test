package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteContentStore struct {
	db *sql.DB
}

// NewContentStore creates a new SQLiteContentStore instance
func NewSQLiteContentStore(dbPath string) (*SQLiteContentStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &SQLiteContentStore{db: db}, nil
}

// Close closes the database connection
func (cs *SQLiteContentStore) Close() error {
	return cs.db.Close()
}

// Create inserts a new content record
func (cs *SQLiteContentStore) Create(content *Content) error {
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
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err = cs.db.Exec(query,
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
func (cs *SQLiteContentStore) GetByID(id string) (*Content, error) {
	query := `
		SELECT id, title, body, author, status, data, created_at, updated_at
		FROM content WHERE id = ?
	`

	var content Content
	var dataJSON []byte

	err := cs.db.QueryRow(query, id).Scan(
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
		if err == sql.ErrNoRows {
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
func (cs *SQLiteContentStore) List() ([]*Content, error) {
	query := `
		SELECT id, title, body, author, status, data, created_at, updated_at
		FROM content ORDER BY created_at DESC LIMIT 100
	`

	rows, err := cs.db.Query(query)
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
func (cs *SQLiteContentStore) Update(content *Content) error {
	// Serialize data to JSON
	dataJSON, err := json.Marshal(content.Data)
	if err != nil {
		return fmt.Errorf("failed to marshal data: %w", err)
	}

	content.UpdatedAt = time.Now()

	query := `
		UPDATE content 
		SET title = ?, body = ?, author = ?, status = ?, data = ?, updated_at = ?
		WHERE id = ?
	`

	result, err := cs.db.Exec(query,
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

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("content not found")
	}

	return nil
}

// Delete removes a content record by ID
func (cs *SQLiteContentStore) Delete(id string) error {
	query := `DELETE FROM content WHERE id = ?`

	result, err := cs.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete content: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("content not found")
	}

	return nil
}
