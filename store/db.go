package store

import (
	"database/sql"
	"flowmesh/config"

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
