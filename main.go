package main

import (
	"api-server/internal/db"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.InitClient()
	InitRouter()

	log.Println("Server starting on port " + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

func init() {
	err := os.Setenv("GOOGLE_PROJECT_ID", "github-blog-b7f62")
	if err != nil {
		log.Fatal("Failed to set FIRESTORE_PROJECT_ID environment variable")
	}
}
