package main

import (
	"backend/internal/handlers"
	"backend/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/cors"
)

var startTime = time.Now()

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
	exerciseRepo := repository.NewExerciseRepository(db)
	exerciseHandler := handlers.NewExerciseHandler(exerciseRepo)
	setRepo := repository.NewSetRepository(db)
	setHandler := handlers.NewSetHandler(setRepo)

	// Router
	r := chi.NewRouter()

	// Apply CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://167.172.234.41"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	r.Use(c.Handler)

	r.Route("/api", func(r chi.Router) {
		r.Route("/health", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				HealthHandler(w, r, db)
			})
		})
		// Sessions collection
		r.Route("/sessions", func(r chi.Router) {
			r.Get("/", sessionHandler.GetSessions)
			r.Post("/", sessionHandler.CreateSession)

			// Single session
			r.Route("/{sessionID}", func(r chi.Router) {
				r.Put("/", sessionHandler.UpdateSession)
				r.Delete("/", sessionHandler.DeleteSession)

				// Exercises for a session
				r.Get("/exercises", exerciseHandler.GetExercises)
				r.Post("/exercises", exerciseHandler.CreateExercise)
			})
		})

		// Single exercise (not tied to session list)
		r.Route("/exercises/{exerciseID}", func(r chi.Router) {
			r.Put("/", exerciseHandler.UpdateExercise)
			r.Delete("/", exerciseHandler.DeleteExercise)

			// Sets for an exercise
			r.Get("/sets", setHandler.GetSets)
			r.Post("/sets", setHandler.CreateSet)
		})

		// Single set
		r.Route("/sets/{setID}", func(r chi.Router) {
			r.Put("/", setHandler.UpdateSet)
			r.Delete("/", setHandler.DeleteSet)
		})
	})

	fmt.Println("Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func HealthHandler(w http.ResponseWriter, r *http.Request, db *pgxpool.Pool) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	// Simple DB check
	dbStatus := "ok"
	if err := db.Ping(ctx); err != nil {
		dbStatus = "unreachable"
	}

	health := map[string]string{
		"status":    "ok",
		"uptime":    time.Since(startTime).String(),
		"version":   "1.0.0", // set via build flag later
		"database":  dbStatus,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(health)
}
