package worker

import (
	"database/sql"
	"flowmesh/model"
	"flowmesh/queue"
	"flowmesh/store"
	"time"

	"github.com/redis/go-redis/v9"
)

func Start(db *sql.DB, redisClient *redis.Client, numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		go executeWorker(db, redisClient)
	}
}

func executeWorker(db *sql.DB, redisClient *redis.Client) {
	for {
		jobID, err := queue.DequeueJob(redisClient)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		job, err := store.GetJob(db, jobID)
		if err != nil {
			continue
		}

		println("Executing job:", job.Name)

		start := time.Now()

		result := model.JobResult{
			JobID:  jobID,
			Status: "success",
			Output: "Job executed",
		}

		result.Duration = int(time.Since(start).Seconds())

		err = store.SaveJobResult(db, result)
		if err != nil {
			continue
		}

		store.UpdateJobStatus(db, jobID, "completed")
	}
}
