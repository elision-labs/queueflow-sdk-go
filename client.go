// client_optimized.go - Performance-optimized extensions for QueueFlow SDK
package queueflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// OptimizedClient extends the base Client with performance enhancements
type OptimizedClient struct {
	*Client
	metrics       MetricsCollector
	logger        Logger
	circuitBreaker CircuitBreaker
	bufferPool    *sync.Pool
}

// MetricsCollector interface for pluggable metrics
type MetricsCollector interface {
	IncrementCounter(name string, tags map[string]string)
	RecordHistogram(name string, value float64, tags map[string]string)
	RecordGauge(name string, value float64, tags map[string]string)
}

// Logger interface for structured logging
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
}

// Field represents a structured log field
type Field struct {
	Key   string
	Value interface{}
}

// CircuitBreaker interface for fault tolerance
type CircuitBreaker interface {
	Execute(func() error) error
	GetState() CircuitState
	Reset()
}

// CircuitState represents the state of a circuit breaker
type CircuitState string

const (
	CircuitClosed    CircuitState = "closed"
	CircuitOpen      CircuitState = "open"
	CircuitHalfOpen  CircuitState = "half_open"
)

// OptimizedClientConfig extends ClientConfig with performance options
type OptimizedClientConfig struct {
	*ClientConfig
	MaxIdleConns        int
	MaxConnsPerHost     int
	IdleConnTimeout     time.Duration
	DisableKeepAlives   bool
	DisableCompression  bool
	EnableMetrics       bool
	EnableCircuitBreaker bool
	CircuitBreakerConfig *CircuitBreakerConfig
}

// NewOptimizedClientConfig creates a new optimized configuration
func NewOptimizedClientConfig() *OptimizedClientConfig {
	return &OptimizedClientConfig{
		ClientConfig:         NewClientConfig(),
		MaxIdleConns:        100,
		MaxConnsPerHost:     10,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		DisableCompression:  false,
		EnableMetrics:       true,
		EnableCircuitBreaker: true,
		CircuitBreakerConfig: DefaultCircuitBreakerConfig(),
	}
}

// NewOptimizedClient creates a new optimized QueueFlow client
func NewOptimizedClient(baseURL, apiKey string, config *OptimizedClientConfig) (*OptimizedClient, error) {
	// Create optimized HTTP client
	transport := &http.Transport{
		MaxIdleConns:        config.MaxIdleConns,
		MaxIdleConnsPerHost: config.MaxConnsPerHost,
		IdleConnTimeout:     config.IdleConnTimeout,
		DisableKeepAlives:   config.DisableKeepAlives,
		DisableCompression:  config.DisableCompression,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	httpClient := &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
	}

	// Create base client
	baseClient := &Client{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: httpClient,
		config:     config.ClientConfig,
	}

	// Create optimized client
	oc := &OptimizedClient{
		Client: baseClient,
		bufferPool: &sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}

	// Initialize circuit breaker if enabled
	if config.EnableCircuitBreaker {
		oc.circuitBreaker = NewSimpleCircuitBreaker(config.CircuitBreakerConfig)
	}

	return oc, nil
}

// SetMetricsCollector sets the metrics collector
func (oc *OptimizedClient) SetMetricsCollector(metrics MetricsCollector) {
	oc.metrics = metrics
}

// SetLogger sets the logger
func (oc *OptimizedClient) SetLogger(logger Logger) {
	oc.logger = logger
}

// CreateJobOptimized creates a job with performance optimizations
func (oc *OptimizedClient) CreateJobOptimized(ctx context.Context, taskName string, payload map[string]interface{}) (string, error) {
	start := time.Now()
	
	// Log the operation
	if oc.logger != nil {
		oc.logger.Debug("creating job",
			Field{"task", taskName},
			Field{"payload_size", len(payload)})
	}

	// Execute with circuit breaker if available
	var jobID string
	var err error

	executeFunc := func() error {
		jobID, err = oc.CreateJob(ctx, taskName, payload)
		return err
	}

	if oc.circuitBreaker != nil {
		err = oc.circuitBreaker.Execute(executeFunc)
	} else {
		err = executeFunc()
	}

	// Record metrics
	if oc.metrics != nil {
		duration := time.Since(start).Seconds()
		tags := map[string]string{
			"task":    taskName,
			"success": fmt.Sprintf("%t", err == nil),
		}
		
		oc.metrics.RecordHistogram("queueflow.job.create.duration", duration, tags)
		oc.metrics.IncrementCounter("queueflow.job.create.total", tags)
		
		if err != nil {
			oc.metrics.IncrementCounter("queueflow.job.create.errors", tags)
		}
	}

	// Log result
	if oc.logger != nil {
		if err != nil {
			oc.logger.Error("failed to create job",
				Field{"task", taskName},
				Field{"error", err.Error()},
				Field{"duration_ms", time.Since(start).Milliseconds()})
		} else {
			oc.logger.Info("job created successfully",
				Field{"task", taskName},
				Field{"job_id", jobID},
				Field{"duration_ms", time.Since(start).Milliseconds()})
		}
	}

	return jobID, err
}

// StreamJobs streams jobs as they are retrieved
func (oc *OptimizedClient) StreamJobs(ctx context.Context, status JobStatus) (<-chan *Job, <-chan error) {
	jobsChan := make(chan *Job, 100) // Buffered channel for better performance
	errChan := make(chan error, 1)
	
	go func() {
		defer close(jobsChan)
		defer close(errChan)
		
		offset := 0
		limit := 100
		
		for {
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
			}
			
			// Log streaming progress
			if oc.logger != nil {
				oc.logger.Debug("streaming jobs batch",
					Field{"offset", offset},
					Field{"limit", limit},
					Field{"status", string(status)})
			}
			
			resp, err := oc.ListJobs(ctx, limit, offset, status)
			if err != nil {
				if oc.logger != nil {
					oc.logger.Error("failed to list jobs",
						Field{"error", err.Error()},
						Field{"offset", offset})
				}
				errChan <- err
				return
			}
			
			// Record metrics for batch
			if oc.metrics != nil {
				oc.metrics.RecordHistogram("queueflow.stream.batch_size", 
					float64(len(resp.Items)), 
					map[string]string{"status": string(status)})
			}
			
			for _, job := range resp.Items {
				jobCopy := job // Important: copy to avoid loop variable issues
				select {
				case jobsChan <- &jobCopy:
				case <-ctx.Done():
					errChan <- ctx.Err()
					return
				}
			}
			
			if !resp.HasMore {
				if oc.logger != nil {
					oc.logger.Info("finished streaming jobs",
						Field{"total_streamed", offset + len(resp.Items)},
						Field{"status", string(status)})
				}
				break
			}
			
			offset += limit
		}
	}()
	
	return jobsChan, errChan
}

// CircuitBreakerConfig configuration for circuit breaker
type CircuitBreakerConfig struct {
	MaxFailures  int
	ResetTimeout time.Duration
	HalfOpenMax  int
}

// DefaultCircuitBreakerConfig returns default circuit breaker configuration
func DefaultCircuitBreakerConfig() *CircuitBreakerConfig {
	return &CircuitBreakerConfig{
		MaxFailures:  5,
		ResetTimeout: 60 * time.Second,
		HalfOpenMax:  3,
	}
}

// SimpleCircuitBreaker implements a basic circuit breaker pattern
type SimpleCircuitBreaker struct {
	config       *CircuitBreakerConfig
	failures     int32
	lastFailTime int64 // Unix timestamp
	halfOpenReqs int32
	state        int32 // 0: closed, 1: open, 2: half-open
	mu           sync.RWMutex
}

// NewSimpleCircuitBreaker creates a new circuit breaker
func NewSimpleCircuitBreaker(config *CircuitBreakerConfig) *SimpleCircuitBreaker {
	return &SimpleCircuitBreaker{
		config: config,
		state:  0, // Start closed
	}
}

// Execute runs a function with circuit breaker protection
func (cb *SimpleCircuitBreaker) Execute(fn func() error) error {
	state := cb.GetState()
	
	switch state {
	case CircuitOpen:
		return &Error{
			Code:    "CIRCUIT_OPEN",
			Message: "circuit breaker is open",
		}
	
	case CircuitHalfOpen:
		// Allow limited requests in half-open state
		reqs := atomic.AddInt32(&cb.halfOpenReqs, 1)
		if reqs > int32(cb.config.HalfOpenMax) {
			atomic.AddInt32(&cb.halfOpenReqs, -1)
			return &Error{
				Code:    "CIRCUIT_HALF_OPEN_LIMIT",
				Message: "circuit breaker half-open limit reached",
			}
		}
		defer atomic.AddInt32(&cb.halfOpenReqs, -1)
	}
	
	// Execute the function
	err := fn()
	
	if err != nil {
		cb.recordFailure()
	} else {
		cb.recordSuccess()
	}
	
	return err
}

// GetState returns the current circuit breaker state
func (cb *SimpleCircuitBreaker) GetState() CircuitState {
	failures := atomic.LoadInt32(&cb.failures)
	lastFail := atomic.LoadInt64(&cb.lastFailTime)
	
	// Check if we should transition from open to half-open
	if failures >= int32(cb.config.MaxFailures) {
		if time.Since(time.Unix(lastFail, 0)) > cb.config.ResetTimeout {
			atomic.StoreInt32(&cb.state, 2) // Half-open
			return CircuitHalfOpen
		}
		atomic.StoreInt32(&cb.state, 1) // Open
		return CircuitOpen
	}
	
	atomic.StoreInt32(&cb.state, 0) // Closed
	return CircuitClosed
}

// Reset manually resets the circuit breaker
func (cb *SimpleCircuitBreaker) Reset() {
	atomic.StoreInt32(&cb.failures, 0)
	atomic.StoreInt32(&cb.state, 0)
	atomic.StoreInt32(&cb.halfOpenReqs, 0)
}

func (cb *SimpleCircuitBreaker) recordFailure() {
	failures := atomic.AddInt32(&cb.failures, 1)
	atomic.StoreInt64(&cb.lastFailTime, time.Now().Unix())
	
	// If we just hit the threshold, transition to open
	if failures == int32(cb.config.MaxFailures) {
		atomic.StoreInt32(&cb.state, 1)
	}
}

func (cb *SimpleCircuitBreaker) recordSuccess() {
	state := atomic.LoadInt32(&cb.state)
	
	// If we're in half-open state and request succeeded, close the circuit
	if state == 2 {
		cb.Reset()
	}
}

// BulkOperationProgress tracks progress of bulk operations
type BulkOperationProgress struct {
	Total       int
	Completed   int32
	Failed      int32
	StartTime   time.Time
	EndTime     *time.Time
}

// GetProgress returns current progress percentage
func (p *BulkOperationProgress) GetProgress() float64 {
	completed := atomic.LoadInt32(&p.Completed)
	failed := atomic.LoadInt32(&p.Failed)
	processed := completed + failed
	
	if p.Total == 0 {
		return 100.0
	}
	
	return float64(processed) / float64(p.Total) * 100
}

// GetDuration returns the duration of the operation
func (p *BulkOperationProgress) GetDuration() time.Duration {
	if p.EndTime != nil {
		return p.EndTime.Sub(p.StartTime)
	}
	return time.Since(p.StartTime)
}

// GetRate returns the processing rate (items per second)
func (p *BulkOperationProgress) GetRate() float64 {
	duration := p.GetDuration().Seconds()
	if duration == 0 {
		return 0
	}
	
	completed := atomic.LoadInt32(&p.Completed)
	failed := atomic.LoadInt32(&p.Failed)
	processed := float64(completed + failed)
	
	return processed / duration
}

// CreateJobsBatchWithProgress creates jobs in batch with progress tracking
func (oc *OptimizedClient) CreateJobsBatchWithProgress(
	ctx context.Context,
	jobs []*BatchJobRequest,
	progressChan chan<- *BulkOperationProgress,
) ([]string, error) {
	progress := &BulkOperationProgress{
		Total:     len(jobs),
		StartTime: time.Now(),
	}
	
	// Send initial progress
	select {
	case progressChan <- progress:
	default:
	}
	
	// Process in parallel chunks for better performance
	chunkSize := 25
	numWorkers := 4
	
	type result struct {
		jobIDs []string
		err    error
		start  int
	}
	
	jobsChan := make(chan struct{ jobs []*BatchJobRequest; start int }, numWorkers)
	resultsChan := make(chan result, len(jobs)/chunkSize+1)
	
	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for chunk := range jobsChan {
				jobIDs, err := oc.CreateJobsBatch(ctx, chunk.jobs)
				
				if err != nil {
					atomic.AddInt32(&progress.Failed, int32(len(chunk.jobs)))
				} else {
					atomic.AddInt32(&progress.Completed, int32(len(chunk.jobs)))
				}
				
				resultsChan <- result{
					jobIDs: jobIDs,
					err:    err,
					start:  chunk.start,
				}
				
				// Send progress update
				select {
				case progressChan <- progress:
				default:
				}
			}
		}()
	}
	
	// Send chunks to workers
	go func() {
		for i := 0; i < len(jobs); i += chunkSize {
			end := i + chunkSize
			if end > len(jobs) {
				end = len(jobs)
			}
			
			select {
			case jobsChan <- struct{ jobs []*BatchJobRequest; start int }{
				jobs:  jobs[i:end],
				start: i,
			}:
			case <-ctx.Done():
				close(jobsChan)
				return
			}
		}
		close(jobsChan)
	}()
	
	// Wait for all workers to complete
	go func() {
		wg.Wait()
		close(resultsChan)
		
		endTime := time.Now()
		progress.EndTime = &endTime
		
		// Send final progress
		select {
		case progressChan <- progress:
		default:
		}
	}()
	
	// Collect results in order
	results := make(map[int][]string)
	var firstError error
	
	for res := range resultsChan {
		if res.err != nil && firstError == nil {
			firstError = res.err
		}
		if res.jobIDs != nil {
			results[res.start] = res.jobIDs
		}
	}
	
	// Reconstruct ordered job IDs
	var allJobIDs []string
	for i := 0; i < len(jobs); i += chunkSize {
		if jobIDs, ok := results[i]; ok {
			allJobIDs = append(allJobIDs, jobIDs...)
		}
	}
	
	// Log final stats
	if oc.logger != nil {
		oc.logger.Info("batch operation completed",
			Field{"total", progress.Total},
			Field{"completed", atomic.LoadInt32(&progress.Completed)},
			Field{"failed", atomic.LoadInt32(&progress.Failed)},
			Field{"duration_ms", progress.GetDuration().Milliseconds()},
			Field{"rate_per_sec", progress.GetRate()})
	}
	
	return allJobIDs, firstError
}