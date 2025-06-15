package main

import (
	"log"
	"todo_api/pkg/handlers"
	"todo_api/server"
)

func main() {
	// Initialize handler
	handler := handlers.NewHandler()
	
	// Initialize routes
	router := handler.InitRoutes()
	
	// Create and start server
	srv := server.New()
	log.Println("Starting server on port 8080...")
	
	if err := srv.Run("8080", router); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
