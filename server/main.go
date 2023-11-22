package main

import (
	"log"
	"server/internal/adapters/api/routes"
)

func main() {
	db, err := routes.Run()
	if err != nil {
		log.Fatalf("Error initializing routes: %v", err)
	}

	routes.Injection(db)
}
