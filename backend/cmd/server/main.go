package main

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

func main() {
	ctx := context.Background()

	// Setup DB
	db, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	// Init repos + handlers
	sessionRepo := repository.NewSessionRepository(db)
	sessionHandler := handlers.NewSessionHandler(sessionRepo)

	// Router
	r := chi.NewRouter()

	// Apply CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://167.172.234.41:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	r.Use(c.Handler)

	// Sessions collection
	r.Route("/sessions", func(r chi.Router) {
		r.Get("/", sessionHandler.GetSessions)
		r.Post("/", sessionHandler.CreateSession)
	})

	// Single session
	r.Route("/session/{id}", func(r chi.Router) {
		r.Put("/", sessionHandler.UpdateSession)
		r.Delete("/", sessionHandler.DeleteSession)
	})

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}
