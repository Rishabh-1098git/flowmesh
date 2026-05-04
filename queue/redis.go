package queue

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func EnqueueJob(client *redis.Client, jobID int) error {
	ctx := context.Background()
	err := client.LPush(ctx, "jobs_queue", jobID).Err()
	return err
}

func DequeueJob(client *redis.Client) (int, error) {
	ctx := context.Background()
	val, err := client.RPop(ctx, "jobs_queue").Result()
	if err != nil {
		return 0, err
	}
	jobID, err := strconv.Atoi(val)
	if err != nil {
		return 0, err
	}
	return jobID, nil
}
