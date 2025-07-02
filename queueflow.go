// Package queueflow provides a Go SDK for the QueueFlow distributed job queue system
package queueflow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultTimeout = 30 * time.Second
	apiVersion     = "/api/v1"
)

// JobStatus represents the status of a job
type JobStatus string

const (
	JobStatusPending   JobStatus = "pending"
	JobStatusRunning   JobStatus = "running"
	JobStatusCompleted JobStatus = "completed"
	JobStatusFailed    JobStatus = "failed"
	JobStatusRetrying  JobStatus = "retrying"
	JobStatusCancelled JobStatus = "cancelled"
)

// WorkflowStatus represents the status of a workflow
type WorkflowStatus string

const (
	WorkflowStatusCreated         WorkflowStatus = "created"
	WorkflowStatusRunning         WorkflowStatus = "running"
	WorkflowStatusCompleted       WorkflowStatus = "completed"
	WorkflowStatusFailed          WorkflowStatus = "failed"
	WorkflowStatusCancelled       WorkflowStatus = "cancelled"
	WorkflowStatusPartiallyFailed WorkflowStatus = "partially_failed"
)

// Client is the main QueueFlow client
type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new QueueFlow client
func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
}

// SetHTTPClient allows setting a custom HTTP client
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

// Request types matching the engine API

// CreateJobRequest represents a job creation request
type CreateJobRequest struct {
	TaskName string                 `json:"task_name"`
	Payload  map[string]interface{} `json:"payload"`
	Config   *JobConfig             `json:"config,omitempty"`
}

// JobConfig represents job configuration
type JobConfig struct {
	Priority   *int    `json:"priority,omitempty"`
	MaxRetries *int    `json:"max_retries,omitempty"`
	Timeout    *int    `json:"timeout,omitempty"` // in seconds
	Queue      *string `json:"queue,omitempty"`
}

// CreateJobResponse represents a job creation response
type CreateJobResponse struct {
	JobID string `json:"job_id"`
}

// Job represents a job in the system
type Job struct {
	ID             string                 `json:"id"`
	QueueName      string                 `json:"queue_name"`
	TaskName       string                 `json:"task_name"`
	Payload        map[string]interface{} `json:"payload"`
	Status         JobStatus              `json:"status"`
	CreatedAt      time.Time              `json:"created_at"`
	StartedAt      *time.Time             `json:"started_at,omitempty"`
	CompletedAt    *time.Time             `json:"completed_at,omitempty"`
	ErrorMessage   string                 `json:"error_message,omitempty"`
	RetryCount     int                    `json:"retry_count"`
	WorkflowID     string                 `json:"workflow_id,omitempty"`
	WorkflowStepID string                 `json:"workflow_step_id,omitempty"`
	Result         map[string]interface{} `json:"result,omitempty"`
	Metadata       map[string]interface{} `json:"metadata"`
}

// BatchJobRequest represents a batch job creation request
type BatchJobRequest struct {
	Jobs []CreateJobRequest `json:"jobs"`
}

// BatchJobResponse represents a batch job creation response
type BatchJobResponse struct {
	JobIDs []string `json:"job_ids"`
	Count  int      `json:"count"`
}

// WorkflowStep represents a step in a workflow
type WorkflowStep struct {
	Name      string                 `json:"name"`
	TaskName  string                 `json:"task_name"`
	Payload   map[string]interface{} `json:"payload"`
	DependsOn []string               `json:"depends_on,omitempty"`
	Config    *JobConfig             `json:"config,omitempty"`
}

// CreateWorkflowRequest represents a workflow creation request
type CreateWorkflowRequest struct {
	Name  string         `json:"name"`
	Steps []WorkflowStep `json:"steps"`
}

// CreateWorkflowResponse represents a workflow creation response
type CreateWorkflowResponse struct {
	WorkflowID string `json:"workflow_id"`
}

// Workflow represents a workflow in the system
type Workflow struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Status      WorkflowStatus         `json:"status"`
	CreatedAt   time.Time              `json:"created_at"`
	StartedAt   *time.Time             `json:"started_at,omitempty"`
	CompletedAt *time.Time             `json:"completed_at,omitempty"`
	Context     map[string]interface{} `json:"context"`
	Metadata    map[string]interface{} `json:"metadata"`
}

// ListJobsResponse represents a paginated list of jobs
type ListJobsResponse struct {
	Jobs    []Job `json:"jobs"`
	Total   int   `json:"total"`
	Limit   int   `json:"limit"`
	Offset  int   `json:"offset"`
	HasMore bool  `json:"has_more"`
}

// ListWorkflowsResponse represents a paginated list of workflows
type ListWorkflowsResponse struct {
	Workflows []Workflow `json:"workflows"`
	Total     int        `json:"total"`
	Limit     int        `json:"limit"`
	Offset    int        `json:"offset"`
	HasMore   bool       `json:"has_more"`
}

// ErrorResponse represents an API error response
type ErrorResponse struct {
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

// Client methods

// doRequest performs an HTTP request
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := c.baseURL + apiVersion + path

	var reqBody io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(data)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	return resp, nil
}

// decodeResponse decodes the response body
func decodeResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errResp ErrorResponse
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("HTTP %d: failed to decode error response", resp.StatusCode)
		}
		return fmt.Errorf("HTTP %d: %s", resp.StatusCode, errResp.Error)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

// CreateJob creates a new job
func (c *Client) CreateJob(ctx context.Context, taskName string, payload map[string]interface{}, config *JobConfig) (string, error) {
	req := CreateJobRequest{
		TaskName: taskName,
		Payload:  payload,
		Config:   config,
	}

	resp, err := c.doRequest(ctx, http.MethodPost, "/jobs", req)
	if err != nil {
		return "", err
	}

	var result CreateJobResponse
	if err := decodeResponse(resp, &result); err != nil {
		return "", err
	}

	return result.JobID, nil
}

// GetJob retrieves a job by ID
func (c *Client) GetJob(ctx context.Context, jobID string) (*Job, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/jobs/"+url.PathEscape(jobID), nil)
	if err != nil {
		return nil, err
	}

	var job Job
	if err := decodeResponse(resp, &job); err != nil {
		return nil, err
	}

	return &job, nil
}

// CancelJob cancels a job
func (c *Client) CancelJob(ctx context.Context, jobID string) error {
	resp, err := c.doRequest(ctx, http.MethodPost, "/jobs/"+url.PathEscape(jobID)+"/cancel", nil)
	if err != nil {
		return err
	}

	return decodeResponse(resp, nil)
}

// ListJobs lists jobs with optional filters
func (c *Client) ListJobs(ctx context.Context, status string, queue string, limit, offset int) (*ListJobsResponse, error) {
	query := url.Values{}
	if status != "" {
		query.Set("status", status)
	}
	if queue != "" {
		query.Set("queue", queue)
	}
	if limit > 0 {
		query.Set("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		query.Set("offset", fmt.Sprintf("%d", offset))
	}

	path := "/jobs"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var result ListJobsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateBatchJobs creates multiple jobs in a single request
func (c *Client) CreateBatchJobs(ctx context.Context, jobs []CreateJobRequest) ([]string, error) {
	req := BatchJobRequest{
		Jobs: jobs,
	}

	resp, err := c.doRequest(ctx, http.MethodPost, "/jobs/batch", req)
	if err != nil {
		return nil, err
	}

	var result BatchJobResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.JobIDs, nil
}

// CreateWorkflow creates a new workflow
func (c *Client) CreateWorkflow(ctx context.Context, name string, steps []WorkflowStep) (string, error) {
	req := CreateWorkflowRequest{
		Name:  name,
		Steps: steps,
	}

	resp, err := c.doRequest(ctx, http.MethodPost, "/workflows", req)
	if err != nil {
		return "", err
	}

	var result CreateWorkflowResponse
	if err := decodeResponse(resp, &result); err != nil {
		return "", err
	}

	return result.WorkflowID, nil
}

// GetWorkflow retrieves a workflow by ID
func (c *Client) GetWorkflow(ctx context.Context, workflowID string) (*Workflow, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, "/workflows/"+url.PathEscape(workflowID), nil)
	if err != nil {
		return nil, err
	}

	var workflow Workflow
	if err := decodeResponse(resp, &workflow); err != nil {
		return nil, err
	}

	return &workflow, nil
}

// CancelWorkflow cancels a workflow
func (c *Client) CancelWorkflow(ctx context.Context, workflowID string) error {
	resp, err := c.doRequest(ctx, http.MethodPost, "/workflows/"+url.PathEscape(workflowID)+"/cancel", nil)
	if err != nil {
		return err
	}

	return decodeResponse(resp, nil)
}

// ListWorkflows lists workflows with optional filters
func (c *Client) ListWorkflows(ctx context.Context, status string, limit, offset int) (*ListWorkflowsResponse, error) {
	query := url.Values{}
	if status != "" {
		query.Set("status", status)
	}
	if limit > 0 {
		query.Set("limit", fmt.Sprintf("%d", limit))
	}
	if offset > 0 {
		query.Set("offset", fmt.Sprintf("%d", offset))
	}

	path := "/workflows"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}

	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var result ListWorkflowsResponse
	if err := decodeResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// WaitForJob polls for job completion
func (c *Client) WaitForJob(ctx context.Context, jobID string, pollInterval time.Duration) (*Job, error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			job, err := c.GetJob(ctx, jobID)
			if err != nil {
				return nil, err
			}

			switch job.Status {
			case JobStatusCompleted, JobStatusFailed, JobStatusCancelled:
				return job, nil
			}
		}
	}
}

// WaitForWorkflow polls for workflow completion
func (c *Client) WaitForWorkflow(ctx context.Context, workflowID string, pollInterval time.Duration) (*Workflow, error) {
	ticker := time.NewTicker(pollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			workflow, err := c.GetWorkflow(ctx, workflowID)
			if err != nil {
				return nil, err
			}

			switch workflow.Status {
			case WorkflowStatusCompleted, WorkflowStatusFailed, WorkflowStatusCancelled, WorkflowStatusPartiallyFailed:
				return workflow, nil
			}
		}
	}
}

// Helper functions

// IntPtr returns a pointer to an int
func IntPtr(i int) *int {
	return &i
}

// StringPtr returns a pointer to a string  
func StringPtr(s string) *string {
	return &s
}