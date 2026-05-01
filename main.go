package main

import (
	"flowmesh/store"
	"fmt"
	"log"
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
}
