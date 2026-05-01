package model

import "time"

type JobResult struct {
	ID         int       `json:"id"`
	JobID      int       `json:"job_id"`
	Status     string    `json:"status"`
	Output     string    `json:"output"`
	Error      string    `json:"error"`
	Duration   int       `json:"duration"`
	ExecutedAt time.Time `json:"executed_at"`
}
