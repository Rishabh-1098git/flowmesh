package main

import (
	"flowmesh/handler"
	"flowmesh/store"
	"flowmesh/worker"
	"fmt"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("FlowMesh starting .....")

	db, err := store.Connect()
	if err != nil {
		log.Fatal("DB connect failed", err)
	}
	fmt.Println("DB connect success")

	err = store.CreateTables(db)
	if err != nil {
		log.Fatal("DB create tables failed", err)
	}
	fmt.Println("DB create tables success")

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	go worker.Start(db, redisClient, 4)

	h := handler.NewPipelineHandler(db, redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/pipeline", h.CreatePipeline)

	fmt.Println("SERVER is running on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
