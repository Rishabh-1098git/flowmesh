# 🚀 FlowMesh

A lightweight distributed job processing system built using Go, Redis, and PostgreSQL.

---

# 📌 Overview

FlowMesh is designed to simulate how real-world systems (like Uber, Netflix, etc.) process background jobs asynchronously.

It allows users to:
- Create pipelines (a group of jobs)
- Store jobs in a database
- Push jobs into a Redis queue
- Process jobs using concurrent workers
- Track execution results and status

---

# 🏗️ System Architecture

```
Client (Postman / API)
        ↓
Create Pipeline API
        ↓
PostgreSQL (store pipeline + jobs)
        ↓
Redis Queue (job IDs)
        ↓
Worker Pool (goroutines)
        ↓
Execute Job
        ↓
Save Results (PostgreSQL)
        ↓
Update Job Status
```

---

# ⚙️ Tech Stack

- **Go (Golang)** → Core backend logic
- **PostgreSQL** → Persistent storage
- **Redis** → Queue system
- **Docker** → Local environment setup

---

# 📂 Project Structure

```
flowmesh/
├── config/        # DB configs
├── handler/       # HTTP handlers (to be added)
├── model/         # Structs (Job, Pipeline, Result)
├── queue/         # Redis queue logic
├── store/         # DB operations
├── worker/        # Worker pool implementation
├── main.go        # Entry point
```

---

# 🔄 Core Flow Explained

## 1. Pipeline Creation
- User submits a pipeline via API
- Pipeline is stored in PostgreSQL
- Each job is stored with `pending` status

## 2. Job Queueing
- Each job ID is pushed into Redis queue

## 3. Worker Execution
- Multiple goroutines run as workers
- Each worker:
  - pulls job ID from Redis
  - fetches job from DB
  - executes job (simulated)
  - saves result
  - updates job status

## 4. Result Storage
- Execution results stored in `job_results`
- Job status updated to `completed`

---

# 🧠 Key Concepts Implemented

## ✅ Worker Pool
- Multiple goroutines process jobs concurrently
- Improves throughput and scalability

## ✅ Queue System (Redis)
- Decouples job creation from execution
- Ensures async processing

## ✅ Persistence Layer (PostgreSQL)
- Stores pipelines, jobs, and results
- Enables tracking and retries

## ✅ Fault Tolerance (basic)
- Workers retry if job fetch fails
- Can be extended with retry logic

---

# 🧪 How to Run

## 1. Start Services
```bash
docker-compose up -d
```

## 2. Run Application
```bash
go run main.go
```

## 3. Test Flow
- Create pipeline (API - coming next)
- Workers will auto-process jobs

---

# 📊 Database Schema

## pipelines
- id
- name
- status
- timestamps

## jobs
- id
- pipeline_id
- status
- payload
- retries

## job_results
- id
- job_id
- status
- output
- duration

---

# 🚧 Current Status

✅ Database layer complete  
✅ Worker pool implemented  
✅ Job execution flow working  
⏳ API layer (in progress)  
⏳ Full pipeline → queue integration  

---

# 🔥 Future Improvements

- Retry mechanism with backoff
- Job failure handling
- Pipeline status aggregation
- Kafka integration for scalability
- Distributed workers (multi-instance)
- Monitoring & logging

---

# 💡 What This Project Demonstrates

- Concurrency using goroutines
- Distributed system basics
- Queue-based architecture
- Separation of concerns (queue, worker, DB)
- Real-world backend design patterns

---

# 🏁 Summary

FlowMesh is a simplified version of real-world job orchestration systems. It demonstrates how asynchronous processing, worker pools, and queue-based architectures work together to build scalable backend systems.

---

💻 Built for learning system design + backend engineering.

