package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/seenthis-ab/content-api/config"
)

// TestContent represents the content structure for API calls
type TestContent struct {
	Title  string                 `json:"title"`
	Body   string                 `json:"body"`
	Author string                 `json:"author"`
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}

// APIResponse represents a generic API response structure
type APIResponse map[string]interface{}

// TestResult holds the result of a single test operation
type TestResult struct {
	Operation string
	ID        string
	Status    int
	Duration  time.Duration
	Error     error
}

// SmokeTest performs CRUD operations in parallel
type SmokeTest struct {
	baseURL    string
	httpClient *http.Client
	results    chan TestResult
	wg         sync.WaitGroup
	semaphore  chan struct{} // semaphore for limiting concurrency
}

// NewSmokeTest creates a new smoke test instance
func NewSmokeTest(baseURL string, parallel int) *SmokeTest {
	return &SmokeTest{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		results:   make(chan TestResult, 1000), // Buffer for results
		semaphore: make(chan struct{}, parallel),
	}
}

// acquire and release helpers
func (st *SmokeTest) acquire() { st.semaphore <- struct{}{} }
func (st *SmokeTest) release() { <-st.semaphore }

// createContent creates a new content item
func (st *SmokeTest) createContent(ctx context.Context, id int) {
	st.acquire()
	defer st.release()
	defer st.wg.Done()

	start := time.Now()

	content := TestContent{
		Title:  fmt.Sprintf("Smoke Test Content %d", id),
		Body:   fmt.Sprintf("This is smoke test content number %d", id),
		Author: "Smoke Tester",
		Status: "draft",
		Data: map[string]interface{}{
			"test_id":    id,
			"created_at": time.Now().Unix(),
		},
	}

	jsonData, err := json.Marshal(content)
	if err != nil {
		st.results <- TestResult{
			Operation: "CREATE",
			ID:        fmt.Sprintf("%d", id),
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to marshal content: %w", err),
		}
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", st.baseURL+"/content", bytes.NewBuffer(jsonData))
	if err != nil {
		st.results <- TestResult{
			Operation: "CREATE",
			ID:        fmt.Sprintf("%d", id),
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to create request: %w", err),
		}
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := st.httpClient.Do(req)
	if err != nil {
		st.results <- TestResult{
			Operation: "CREATE",
			ID:        fmt.Sprintf("%d", id),
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to execute request: %w", err),
		}
		return
	}
	defer resp.Body.Close()

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		st.results <- TestResult{
			Operation: "CREATE",
			ID:        fmt.Sprintf("%d", id),
			Status:    resp.StatusCode,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to decode response: %w", err),
		}
		return
	}

	st.results <- TestResult{
		Operation: "CREATE",
		ID:        apiResp["id"].(string), // Assuming ID is always a string in the new APIResponse
		Status:    resp.StatusCode,
		Duration:  time.Since(start),
		Error:     nil,
	}

	// If creation was successful, perform READ, UPDATE, DELETE operations in sequence
	if resp.StatusCode == 200 || resp.StatusCode == 201 {
		st.wg.Add(3)
		go func() {
			// READ first
			st.readContent(ctx, apiResp["id"].(string))
			// Small delay to ensure READ completes
			time.Sleep(10 * time.Millisecond)
			// UPDATE second
			st.updateContent(ctx, apiResp["id"].(string))
			// Small delay to ensure UPDATE completes
			time.Sleep(10 * time.Millisecond)
			// DELETE last
			st.deleteContent(ctx, apiResp["id"].(string))
		}()
	}
}

// readContent reads a content item by ID
func (st *SmokeTest) readContent(ctx context.Context, id string) {
	st.acquire()
	defer st.release()
	defer st.wg.Done()

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", st.baseURL+"/content/"+id, nil)
	if err != nil {
		st.results <- TestResult{
			Operation: "READ",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to create request: %w", err),
		}
		return
	}

	resp, err := st.httpClient.Do(req)
	if err != nil {
		st.results <- TestResult{
			Operation: "READ",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to execute request: %w", err),
		}
		return
	}
	defer resp.Body.Close()

	st.results <- TestResult{
		Operation: "READ",
		ID:        id,
		Status:    resp.StatusCode,
		Duration:  time.Since(start),
		Error:     nil,
	}
}

// updateContent updates a content item
func (st *SmokeTest) updateContent(ctx context.Context, id string) {
	st.acquire()
	defer st.release()
	defer st.wg.Done()

	start := time.Now()

	updateData := map[string]interface{}{
		"title":  fmt.Sprintf("Updated Smoke Test Content %s", id),
		"status": "published",
		"author": "Updated Smoke Tester",
		"body":   "Updated Smoke Test Body",
		"data": map[string]interface{}{
			"updated_at": time.Now().Unix(),
			"updated_by": "smoke_test",
		},
	}

	jsonData, err := json.Marshal(updateData)
	if err != nil {
		st.results <- TestResult{
			Operation: "UPDATE",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to marshal update data: %w", err),
		}
		return
	}

	fmt.Println("Updating content", id, string(jsonData))

	req, err := http.NewRequestWithContext(ctx, "PUT", st.baseURL+"/content/"+id, bytes.NewBuffer(jsonData))
	if err != nil {
		st.results <- TestResult{
			Operation: "UPDATE",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to create request: %w", err),
		}
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := st.httpClient.Do(req)
	if err != nil {
		st.results <- TestResult{
			Operation: "UPDATE",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to execute request: %w", err),
		}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errMsg string
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		errMsg = buf.String()
		log.Printf("UPDATE error response for id %s: %s", id, errMsg)
	}

	st.results <- TestResult{
		Operation: "UPDATE",
		ID:        id,
		Status:    resp.StatusCode,
		Duration:  time.Since(start),
		Error:     nil,
	}
}

// deleteContent deletes a content item
func (st *SmokeTest) deleteContent(ctx context.Context, id string) {
	st.acquire()
	defer st.release()
	defer st.wg.Done()

	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "DELETE", st.baseURL+"/content/"+id, nil)
	if err != nil {
		st.results <- TestResult{
			Operation: "DELETE",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to create request: %w", err),
		}
		return
	}

	resp, err := st.httpClient.Do(req)
	if err != nil {
		st.results <- TestResult{
			Operation: "DELETE",
			ID:        id,
			Status:    0,
			Duration:  time.Since(start),
			Error:     fmt.Errorf("failed to execute request: %w", err),
		}
		return
	}
	defer resp.Body.Close()

	st.results <- TestResult{
		Operation: "DELETE",
		ID:        id,
		Status:    resp.StatusCode,
		Duration:  time.Since(start),
		Error:     nil,
	}
}

// Run executes the smoke test
func (st *SmokeTest) Run(numIterations int) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	log.Printf("Starting smoke test with %d iterations...", numIterations)

	// Start result collector
	go func() {
		st.wg.Wait()
		close(st.results)
	}()

	// Start create operations
	for i := 0; i < numIterations; i++ {
		st.wg.Add(1)
		go st.createContent(ctx, i+1)
	}

	// Collect and analyze results
	operationStats := make(map[string]struct {
		count    int
		success  int
		failures int
		totalDur time.Duration
		minDur   time.Duration
		maxDur   time.Duration
	})

	for result := range st.results {
		stats := operationStats[result.Operation]
		stats.count++
		stats.totalDur += result.Duration

		if result.Error != nil || result.Status >= 400 {
			stats.failures++
			log.Printf("FAILURE: %s %s - Status: %d, Error: %v", result.Operation, result.ID, result.Status, result.Error)
		} else {
			stats.success++
		}

		if stats.minDur == 0 || result.Duration < stats.minDur {
			stats.minDur = result.Duration
		}
		if result.Duration > stats.maxDur {
			stats.maxDur = result.Duration
		}

		operationStats[result.Operation] = stats
	}

	// Print results
	log.Printf("\n=== Smoke Test Results ===")
	totalOperations := 0
	totalSuccess := 0
	totalFailures := 0

	for op, stats := range operationStats {
		avgDur := time.Duration(0)
		if stats.count > 0 {
			avgDur = stats.totalDur / time.Duration(stats.count)
		}

		log.Printf("%s: %d total, %d success, %d failures, avg: %v, min: %v, max: %v",
			op, stats.count, stats.success, stats.failures, avgDur, stats.minDur, stats.maxDur)

		totalOperations += stats.count
		totalSuccess += stats.success
		totalFailures += stats.failures
	}

	log.Printf("\nTotal: %d operations, %d success, %d failures", totalOperations, totalSuccess, totalFailures)

	if totalFailures == 0 {
		log.Printf("SUCCESS: All operations completed successfully")
	} else {
		log.Printf("FAILURE: %d operations failed", totalFailures)
	}
}

func main() {
	var (
		baseURL    = flag.String("url", "http://localhost:8888", "Base URL of the API")
		iterations = flag.Int("n", 10, "Number of iterations to run")
		parallel   = flag.Int("parallel", 10, "Maximum number of parallel API calls")
	)
	flag.Parse()

	// Load config to show which database is being used
	dbConfig := config.LoadDatabaseConfig()
	log.Printf("Using database: %s", dbConfig.GetConnectionString())

	smokeTest := NewSmokeTest(*baseURL, *parallel)
	smokeTest.Run(*iterations)
}
