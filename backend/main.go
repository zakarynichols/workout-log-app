package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/rs/cors"
)

// Make conn a package-level variable so yourHandler can access it
var conn *pgx.Conn

func main() {
	ctx := context.Background()
	var err error
	conn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	var greeting string
	err = conn.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Query result:", greeting)

	mux := http.NewServeMux()

	// Register handler
	mux.HandleFunc("/", yourHandler)

	// Setup CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://167.172.234.41:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", handler)
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Query the database for a message
	var message string
	err := conn.QueryRow(ctx, "SELECT 'Hello from DB!'").Scan(&message)
	if err != nil {
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": message,
	}

	// Encode response as JSON and write it
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
