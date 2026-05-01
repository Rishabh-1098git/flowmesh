package model

import "time"

type Job struct {
	ID         int       `json:"id"`
	PipelineID int       `json:"pipeline_id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	Payload    string    `json:"payload"`
	RetryCount int       `json:"retry_count"`
	MaxRetries int       `json:"max_retries"`
	Error      string    `json:"error"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
