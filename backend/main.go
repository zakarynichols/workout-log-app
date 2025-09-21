package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

// Use a connection pool instead of single connection
var db *pgxpool.Pool

func main() {
	ctx := context.Background()

	var err error
	db, err = pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Test query on startup
	var greeting string
	err = db.QueryRow(ctx, "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Startup query failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Query result:", greeting)

	mux := http.NewServeMux()
	mux.HandleFunc("/", yourHandler)

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

	fmt.Printf("Incoming request from %s %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

	// Each request uses the pool to safely get a connection
	var message string
	err := db.QueryRow(ctx, "SELECT 'Hello from DB!'").Scan(&message)
	if err != nil {
		fmt.Printf("DB query error for %s: %v\n", r.RemoteAddr, err)
		http.Error(w, fmt.Sprintf("DB query error: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": message,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("Response encoding error for %s: %v\n", r.RemoteAddr, err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Responded successfully to %s\n", r.RemoteAddr)
}
