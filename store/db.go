package store

import (
	"database/sql"
	"flowmesh/config"
	"flowmesh/model"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Dburl)

	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS pipelines (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'draft',
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS jobs (
			id SERIAL PRIMARY KEY,
			pipeline_id INT NOT NULL REFERENCES pipelines(id) ON DELETE CASCADE,
			name TEXT NOT NULL,
			status TEXT NOT NULL DEFAULT 'pending',
			payload TEXT NOT NULL,
			max_retries INT NOT NULL DEFAULT 3,
			retry_count INT NOT NULL DEFAULT 0,
			error TEXT,
			created_at TIMESTAMP DEFAULT NOW(),
			updated_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS job_results (
			id SERIAL PRIMARY KEY,
			job_id INT NOT NULL REFERENCES jobs(id) ON DELETE CASCADE,
			status TEXT NOT NULL,
			output TEXT,
			error TEXT,
			duration INT NOT NULL,
			executed_at TIMESTAMP DEFAULT NOW()
		)
	`)
	if err != nil {
		return err
	}

	return nil
}

func CreateJob(db *sql.DB, job model.Job) (model.Job, error) {
	err := db.QueryRow(
		`INSERT INTO jobs (pipeline_id, name, status, payload, max_retries) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`, job.PipelineID, job.Name, job.Status, job.Payload, job.MaxRetries,
	).Scan(
		&job.ID,
		&job.CreatedAt,
		&job.UpdatedAt)
	return job, err
}

func GetJob(db *sql.DB, id int) (model.Job, error) {
	var job model.Job

	err := db.QueryRow(
		`SELECT id, pipeline_id, name, status, payload, max_retries, retry_count, created_at, updated_at FROM jobs WHERE id = $1`, id).Scan(
		&job.ID,
		&job.PipelineID,
		&job.Name,
		&job.Status,
		&job.Payload,
		&job.MaxRetries,
		&job.RetryCount,
		&job.CreatedAt,
		&job.UpdatedAt,
	)
	return job, err
}

func SaveJobResult(db *sql.DB, result model.JobResult) error {
	_, err := db.Exec(
		`INSERT INTO job_results (job_id, status, output, error, duration)
VALUES ($1, $2, $3, $4, $5)`,
		result.JobID,
		result.Status,
		result.Output,
		result.Error,
		result.Duration)
	return err
}

func UpdateJobStatus(db *sql.DB, id int, status string) error {
	_, err := db.Exec(
		`UPDATE jobs SET status = $1 WHERE id = $2`, status, id)

	return err
}

func CreatePipeline(db *sql.DB, pipeline model.Pipeline) (model.Pipeline, error) {
	err := db.QueryRow(
		`INSERT INTO pipelines (name, status) VALUES ($1, $2) RETURNING id, created_date, updated_at`, pipeline.Name, pipeline.Status,
	).Scan(
		&pipeline.ID,
		&pipeline.CreatedAt,
		&pipeline.UpdatedAt)
	return pipeline, err
}

func GetPipeline(db *sql.DB, id int) (model.Pipeline, error) {
	var p model.Pipeline

	err := db.QueryRow(
		`SELECT id, name, status, created_at, updated_at
		 FROM pipelines WHERE id = $1`,
		id,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Status,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	return p, err
}
