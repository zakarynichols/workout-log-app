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

// Package-level DB connection
var conn *pgx.Conn

func main() {
	ctx := context.Background()
	var err error

	// Connect to DB
	conn, err = pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	// Test query on startup
	var greeting string
	err = conn.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Query result:", greeting)

	mux := http.NewServeMux()
	mux.HandleFunc("/", yourHandler)

	// Setup CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://167.172.234.41:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func yourHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Log incoming request
	fmt.Printf("Incoming request from %s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

	// Query DB for message
	var message string
	err := conn.QueryRow(ctx, "SELECT 'Hello from DB!'").Scan(&message)
	if err != nil {
		fmt.Printf("DB query error for %s: %v\n", r.RemoteAddr, err)
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	// Write JSON response
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Response encoding error for %s: %v\n", r.RemoteAddr, err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Log successful response
	fmt.Printf("Responded successfully to %s\n", r.RemoteAddr)
}
