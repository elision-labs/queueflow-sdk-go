package main

import (
	"context"
	"fmt"
	"log"
	"time"

	qf "github.com/queueflow/queueflow-go"
)

func main() {
	// Initialize client
	client := qf.NewClient("http://localhost:8080", "your-api-key")

	ctx := context.Background()

	// Example 1: Create a simple job
	jobID, err := client.CreateJob(ctx, "send_email", map[string]interface{}{
		"to":      "user@example.com",
		"subject": "Welcome!",
		"body":    "Thank you for signing up.",
	}, nil)
	if err != nil {
		log.Fatalf("Failed to create job: %v", err)
	}
	fmt.Printf("Created job: %s\n", jobID)

	// Example 2: Create a job with configuration
	jobID2, err := client.CreateJob(ctx, "process_data", map[string]interface{}{
		"file_url": "https://example.com/data.csv",
		"format":   "csv",
	}, &qf.JobConfig{
		Priority:   qf.IntPtr(5),
		MaxRetries: qf.IntPtr(5),
		Timeout:    qf.IntPtr(600), // 10 minutes
		Queue:      qf.StringPtr("high-priority"),
	})
	if err != nil {
		log.Fatalf("Failed to create job with config: %v", err)
	}
	fmt.Printf("Created high-priority job: %s\n", jobID2)

	// Example 3: Create batch jobs
	batchJobs := []qf.CreateJobRequest{
		{
			TaskName: "resize_image",
			Payload: map[string]interface{}{
				"image_url": "https://example.com/image1.jpg",
				"width":     800,
				"height":    600,
			},
		},
		{
			TaskName: "resize_image",
			Payload: map[string]interface{}{
				"image_url": "https://example.com/image2.jpg",
				"width":     800,
				"height":    600,
			},
		},
	}

	jobIDs, err := client.CreateBatchJobs(ctx, batchJobs)
	if err != nil {
		log.Fatalf("Failed to create batch jobs: %v", err)
	}
	fmt.Printf("Created batch jobs: %v\n", jobIDs)

	// Example 4: Create a workflow
	steps := []qf.WorkflowStep{
		{
			Name:     "download",
			TaskName: "download_file",
			Payload: map[string]interface{}{
				"url": "https://example.com/video.mp4",
			},
		},
		{
			Name:      "transcode",
			TaskName:  "transcode_video",
			Payload:   map[string]interface{}{"format": "webm"},
			DependsOn: []string{"download"},
		},
		{
			Name:      "thumbnail",
			TaskName:  "generate_thumbnail",
			Payload:   map[string]interface{}{},
			DependsOn: []string{"download"},
		},
		{
			Name:      "upload",
			TaskName:  "upload_to_cdn",
			Payload:   map[string]interface{}{},
			DependsOn: []string{"transcode", "thumbnail"},
		},
	}

	workflowID, err := client.CreateWorkflow(ctx, "video_processing", steps)
	if err != nil {
		log.Fatalf("Failed to create workflow: %v", err)
	}
	fmt.Printf("Created workflow: %s\n", workflowID)

	// Example 5: Monitor job status
	job, err := client.GetJob(ctx, jobID)
	if err != nil {
		log.Fatalf("Failed to get job: %v", err)
	}
	fmt.Printf("Job status: %s\n", job.Status)

	// Example 6: Wait for job completion
	completedJob, err := client.WaitForJob(ctx, jobID, 2*time.Second)
	if err != nil {
		log.Fatalf("Failed to wait for job: %v", err)
	}
	fmt.Printf("Job completed with status: %s\n", completedJob.Status)

	// Example 7: List jobs
	jobs, err := client.ListJobs(ctx, string(qf.JobStatusPending), "", 10, 0)
	if err != nil {
		log.Fatalf("Failed to list jobs: %v", err)
	}
	fmt.Printf("Found %d pending jobs (total: %d)\n", len(jobs.Jobs), jobs.Total)

	// Example 8: Cancel a job
	if err := client.CancelJob(ctx, jobID2); err != nil {
		log.Printf("Failed to cancel job: %v", err)
	} else {
		fmt.Printf("Successfully cancelled job: %s\n", jobID2)
	}

	// Example 9: Monitor workflow
	workflow, err := client.GetWorkflow(ctx, workflowID)
	if err != nil {
		log.Fatalf("Failed to get workflow: %v", err)
	}
	fmt.Printf("Workflow status: %s\n", workflow.Status)

	// Example 10: Wait for workflow completion
	completedWorkflow, err := client.WaitForWorkflow(ctx, workflowID, 5*time.Second)
	if err != nil {
		log.Fatalf("Failed to wait for workflow: %v", err)
	}
	fmt.Printf("Workflow completed with status: %s\n", completedWorkflow.Status)
}