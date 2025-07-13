package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"crypto/rand"

	"github.com/danielgtaylor/huma/v2"
	"github.com/oklog/ulid/v2"
	"go.uber.org/zap"

	"github.com/seenthis-ab/content-api/config"
	"github.com/seenthis-ab/content-api/models"
)

// ContentStore interface for dependency injection
type ContentStore interface {
	Create(content *models.Content) error
	GetByID(id string) (*models.Content, error)
	List() ([]*models.Content, error)
	Update(content *models.Content) error
	Delete(id string) error
	Close() error
}

// ContentHandlers contains all content-related HTTP handlers
type ContentHandlers struct {
	store ContentStore
}

// NewContentHandlers creates a new ContentHandlers instance
func NewContentHandlers(store ContentStore) *ContentHandlers {
	return &ContentHandlers{
		store: store,
	}
}

// CreateContentInput represents the request body for creating content
type CreateContentInput struct {
	Body struct {
		Title  string                 `json:"title" required:"true"`
		Body   string                 `json:"body" required:"true"`
		Author string                 `json:"author" required:"true"`
		Status string                 `json:"status" enum:"draft,published,archived" default:"draft"`
		Data   map[string]interface{} `json:"data,omitempty"`
	}
}

// CreateContentOutput represents the response for creating content
type CreateContentOutput struct {
	Body models.Content `json:"body"`
}

// CreateContent handles POST /content requests
func (h *ContentHandlers) CreateContent(ctx context.Context, input *CreateContentInput) (*CreateContentOutput, error) {
	// Get logger with request ID from context
	logger := config.GetLoggerWithRequestID(ctx)

	// Generate a ULID for the content ID (lowercase)
	ts := time.Now().UTC()
	entropy := ulid.Monotonic(rand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(ts), entropy).String()
	id = strings.ToLower(id)

	logger.Info("Creating new content",
		zap.String("content_id", id),
		zap.String("title", input.Body.Title),
		zap.String("author", input.Body.Author),
	)

	if input.Body.Data == nil {
		input.Body.Data = map[string]interface{}{}
	}

	content := &models.Content{
		ID:     id,
		Title:  input.Body.Title,
		Body:   input.Body.Body,
		Author: input.Body.Author,
		Status: input.Body.Status,
		Data:   input.Body.Data,
	}

	if err := h.store.Create(content); err != nil {
		logger.Error("Failed to create content",
			zap.String("content_id", id),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to create content: %w", err)
	}

	logger.Info("Successfully created content",
		zap.String("content_id", id),
	)

	return &CreateContentOutput{Body: *content}, nil
}

// GetContentInput represents the request parameters for getting content by ID
type GetContentInput struct {
	ID string `path:"id"`
}

// GetContentOutput represents the response for getting content by ID
type GetContentOutput struct {
	Body models.Content `json:"body"`
}

// GetContent handles GET /content/{id} requests
func (h *ContentHandlers) GetContent(ctx context.Context, input *GetContentInput) (*GetContentOutput, error) {
	// Get logger with request ID from context
	logger := config.GetLoggerWithRequestID(ctx)

	logger.Info("Fetching content by ID",
		zap.String("content_id", input.ID),
	)

	content, err := h.store.GetByID(input.ID)
	if err != nil {
		logger.Warn("Content not found",
			zap.String("content_id", input.ID),
			zap.Error(err),
		)
		return nil, huma.Error404NotFound("Content not found")
	}

	logger.Info("Successfully retrieved content",
		zap.String("content_id", input.ID),
		zap.String("title", content.Title),
	)

	return &GetContentOutput{Body: *content}, nil
}

// ListContentOutput represents the response for listing content
type ListContentOutput struct {
	Body []models.Content `json:"body"`
}

// ListContent handles GET /content requests
func (h *ContentHandlers) ListContent(ctx context.Context, input *struct{}) (*ListContentOutput, error) {
	contents, err := h.store.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list content: %w", err)
	}

	result := make([]models.Content, len(contents))
	for i, c := range contents {
		result[i] = *c
	}

	return &ListContentOutput{Body: result}, nil
}

// UpdateContentInput represents the request body and parameters for updating content
type UpdateContentInput struct {
	ID   string `path:"id"`
	Body struct {
		Title  *string                `json:"title"`
		Body   *string                `json:"body"`
		Author *string                `json:"author"`
		Status *string                `json:"status" enum:"draft,published,archived"`
		Data   map[string]interface{} `json:"data,omitempty"`
	}
}

// UpdateContentOutput represents the response for updating content
type UpdateContentOutput struct {
	Body models.Content `json:"body"`
}

// UpdateContent handles PUT /content/{id} requests
func (h *ContentHandlers) UpdateContent(ctx context.Context, input *UpdateContentInput) (*UpdateContentOutput, error) {
	content, err := h.store.GetByID(input.ID)
	if err != nil {
		return nil, huma.Error404NotFound("Content not found")
	}

	if input.Body.Title != nil {
		content.Title = *input.Body.Title
	}
	if input.Body.Body != nil {
		content.Body = *input.Body.Body
	}
	if input.Body.Author != nil {
		content.Author = *input.Body.Author
	}
	if input.Body.Status != nil {
		content.Status = *input.Body.Status
	}
	if input.Body.Data != nil {
		content.Data = input.Body.Data
	}

	if err := h.store.Update(content); err != nil {
		return nil, fmt.Errorf("failed to update content: %w", err)
	}

	return &UpdateContentOutput{Body: *content}, nil
}

// DeleteContentInput represents the request parameters for deleting content
type DeleteContentInput struct {
	ID string `path:"id"`
}

// DeleteContentOutput represents the response for deleting content
type DeleteContentOutput struct {
	Body struct{} `json:"body"`
}

// DeleteContent handles DELETE /content/{id} requests
func (h *ContentHandlers) DeleteContent(ctx context.Context, input *DeleteContentInput) (*DeleteContentOutput, error) {
	if err := h.store.Delete(input.ID); err != nil {
		return nil, huma.Error404NotFound("Content not found")
	}
	return &DeleteContentOutput{}, nil
}
