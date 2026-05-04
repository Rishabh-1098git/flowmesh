package handler

import (
	"database/sql"
	"encoding/json"
	"flowmesh/model"
	"flowmesh/queue"
	"flowmesh/store"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type PipelineHandler struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewPipelineHandler(db *sql.DB, redisClient *redis.Client) *PipelineHandler {
	return &PipelineHandler{
		db:          db,
		redisClient: redisClient,
	}
}

func (h *PipelineHandler) CreatePipeline(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var pipeline model.Pipeline

	err := json.NewDecoder(r.Body).Decode(&pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pipeline.Status = "pending"

	createdPipeline, err := store.CreatePipeline(h.db, pipeline)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, job := range pipeline.Jobs {
		job.PipelineID = createdPipeline.ID
		job.Status = "pending"

		createdJob, err := store.CreateJob(h.db, job)
		if err != nil {
			println("failed to create job:", err.Error())
			continue
		}

		err = queue.EnqueueJob(h.redisClient, createdJob.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPipeline)
}
