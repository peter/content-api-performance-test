package main

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"github.com/seenthis-ab/content-api/config"
	postgresModels "github.com/seenthis-ab/content-api/postgres-models"
	sqliteModels "github.com/seenthis-ab/content-api/sqlite-models"
)

// ContentStore interface for both backends
type ContentStore interface {
	Create(content *Content) error
	List() ([]*Content, error)
	Close() error
}

// Content is a generic type for both backends
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

// Adapter types to satisfy ContentStore interface for each backend
type sqliteContentStoreAdapter struct{ s *sqliteModels.ContentStore }

func (a *sqliteContentStoreAdapter) Create(content *Content) error {
	return a.s.Create((*sqliteModels.Content)(content))
}
func (a *sqliteContentStoreAdapter) List() ([]*Content, error) {
	list, err := a.s.List()
	if err != nil {
		return nil, err
	}
	res := make([]*Content, len(list))
	for i, c := range list {
		res[i] = (*Content)(c)
	}
	return res, nil
}
func (a *sqliteContentStoreAdapter) Close() error {
	return a.s.Close()
}

type postgresContentStoreAdapter struct{ s *postgresModels.ContentStore }

func (a *postgresContentStoreAdapter) Create(content *Content) error {
	return a.s.Create((*postgresModels.Content)(content))
}
func (a *postgresContentStoreAdapter) List() ([]*Content, error) {
	list, err := a.s.List()
	if err != nil {
		return nil, err
	}
	res := make([]*Content, len(list))
	for i, c := range list {
		res[i] = (*Content)(c)
	}
	return res, nil
}
func (a *postgresContentStoreAdapter) Close() error {
	a.s.Close()
	return nil
}

// Sample data for generating realistic content
var (
	titles = []string{
		"Getting Started with Go Programming",
		"Advanced Database Design Patterns",
		"Microservices Architecture Best Practices",
		"Understanding RESTful APIs",
		"Cloud Computing Fundamentals",
		"DevOps Pipeline Optimization",
		"Machine Learning Basics",
		"Web Security Essentials",
		"Container Orchestration with Kubernetes",
		"API Design Principles",
		"Database Performance Tuning",
		"Software Testing Strategies",
		"Agile Development Methodologies",
		"Continuous Integration Best Practices",
		"System Design Interview Preparation",
		"Data Structures and Algorithms",
		"Network Programming Fundamentals",
		"Mobile App Development Guide",
		"Blockchain Technology Overview",
		"Artificial Intelligence Trends",
	}

	bodies = []string{
		"This comprehensive guide covers all the essential concepts you need to know to get started with Go programming language. From basic syntax to advanced concurrency patterns, this article provides practical examples and best practices.",
		"Database design is a critical aspect of any application. This article explores advanced patterns and techniques for designing scalable, maintainable database schemas that can handle complex business requirements.",
		"Microservices have become the standard architecture for modern applications. Learn about the best practices for designing, implementing, and maintaining microservices-based systems.",
		"RESTful APIs are the backbone of modern web applications. This guide covers the principles, conventions, and best practices for designing effective REST APIs.",
		"Cloud computing has revolutionized how we build and deploy applications. Understand the fundamental concepts, service models, and deployment strategies.",
		"DevOps is more than just tools and automation. This article explores how to build efficient, reliable pipelines that accelerate software delivery.",
		"Machine learning is transforming industries across the globe. Get started with the basics of ML algorithms, data preprocessing, and model evaluation.",
		"Web security is crucial in today's interconnected world. Learn about common vulnerabilities, attack vectors, and defensive strategies.",
		"Kubernetes has become the de facto standard for container orchestration. Master the concepts and practices for deploying and managing containerized applications.",
		"Good API design is essential for developer experience and system maintainability. Discover the principles that make APIs intuitive and efficient.",
	}

	authors = []string{
		"Alice Johnson",
		"Bob Smith",
		"Carol Davis",
		"David Wilson",
		"Eva Brown",
		"Frank Miller",
		"Grace Lee",
		"Henry Taylor",
		"Iris Garcia",
		"Jack Anderson",
		"Kate Martinez",
		"Liam Thompson",
		"Mia Rodriguez",
		"Noah White",
		"Olivia Clark",
		"Paul Lewis",
		"Quinn Hall",
		"Ruby Young",
		"Sam Allen",
		"Tina King",
	}

	statuses = []string{"draft", "published", "archived"}

	sampleData = []map[string]interface{}{
		{"category": "programming", "tags": []string{"go", "tutorial", "beginner"}, "readTime": 15},
		{"category": "database", "tags": []string{"sql", "design", "advanced"}, "readTime": 25},
		{"category": "architecture", "tags": []string{"microservices", "distributed", "scalable"}, "readTime": 30},
		{"category": "api", "tags": []string{"rest", "http", "design"}, "readTime": 20},
		{"category": "cloud", "tags": []string{"aws", "azure", "gcp"}, "readTime": 18},
		{"category": "devops", "tags": []string{"ci/cd", "automation", "pipeline"}, "readTime": 22},
		{"category": "ml", "tags": []string{"ai", "algorithms", "data"}, "readTime": 35},
		{"category": "security", "tags": []string{"web", "vulnerabilities", "defense"}, "readTime": 28},
		{"category": "containers", "tags": []string{"docker", "kubernetes", "orchestration"}, "readTime": 32},
		{"category": "design", "tags": []string{"principles", "best-practices", "patterns"}, "readTime": 16},
	}
)

// randomChoice selects a random element from a slice
func randomChoice[T any](slice []T) T {
	n, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(slice))))
	return slice[n.Int64()]
}

// generateULID generates a lowercase ULID
func generateULID() string {
	ts := time.Now().UTC()
	entropy := ulid.Monotonic(cryptorand.Reader, 0)
	id := ulid.MustNew(ulid.Timestamp(ts), entropy).String()
	return strings.ToLower(id)
}

// generateTestContent creates a single test content item
func generateTestContent() *Content {
	title := randomChoice(titles)
	body := randomChoice(bodies)
	author := randomChoice(authors)
	status := randomChoice(statuses)
	data := randomChoice(sampleData)

	// Add some randomization to make data more varied
	data["views"] = rand.Intn(10000)
	data["likes"] = rand.Intn(1000)
	data["published"] = time.Now().Add(-time.Duration(rand.Intn(365)) * 24 * time.Hour)

	return &Content{
		ID:     generateULID(),
		Title:  title,
		Body:   body,
		Author: author,
		Status: status,
		Data:   data,
	}
}

func main() {
	// Load database configuration
	dbConfig := config.LoadDatabaseConfig()
	var contentStore ContentStore
	var err error

	// Initialize database connection based on engine
	switch config.GetDatabaseEngine() {
	case "postgres":
		fmt.Println("Using Postgres database engine")
		var pgStore *postgresModels.ContentStore
		pgStore, err = postgresModels.NewContentStore(dbConfig.GetConnectionString())
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize Postgres database: %v", err))
		}
		contentStore = &postgresContentStoreAdapter{s: pgStore}
	case "sqlite":
		fallthrough
	default:
		fmt.Println("Using SQLite database engine")
		var sqlStore *sqliteModels.ContentStore
		sqlStore, err = sqliteModels.NewContentStore(dbConfig.GetConnectionString())
		if err != nil {
			panic(fmt.Sprintf("Failed to initialize SQLite database: %v", err))
		}
		contentStore = &sqliteContentStoreAdapter{s: sqlStore}
	}
	defer contentStore.Close()

	// Number of test records to generate
	numRecords := 100000
	fmt.Printf("Generating %d test content records...\n", numRecords)

	// Track progress
	startTime := time.Now()
	batchSize := 100

	for i := 0; i < numRecords; i++ {
		content := generateTestContent()

		if err := contentStore.Create(content); err != nil {
			fmt.Printf("Error creating content %d: %v\n", i+1, err)
			continue
		}

		// Progress reporting
		if (i+1)%batchSize == 0 {
			elapsed := time.Since(startTime)
			rate := float64(i+1) / elapsed.Seconds()
			fmt.Printf("Created %d records (%.1f records/sec)\n", i+1, rate)
		}
	}

	elapsed := time.Since(startTime)
	fmt.Printf("\n✅ Successfully created %d test records in %.2f seconds\n", numRecords, elapsed.Seconds())
	fmt.Printf("Average rate: %.1f records/second\n", float64(numRecords)/elapsed.Seconds())

	// Verify the count
	contents, err := contentStore.List()
	if err != nil {
		fmt.Printf("Warning: Could not verify record count: %v\n", err)
	} else {
		fmt.Printf("Total records in database: %d\n", len(contents))
	}
}
