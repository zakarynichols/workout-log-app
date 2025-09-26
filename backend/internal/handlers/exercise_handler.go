package handlers

import (
	"backend/internal/models"
	"backend/internal/repository"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ExerciseHandler struct {
	repo *repository.ExerciseRepository
}

func NewExerciseHandler(repo *repository.ExerciseRepository) *ExerciseHandler {
	return &ExerciseHandler{repo: repo}
}

// List exercises for a session
func (h *ExerciseHandler) GetExercises(w http.ResponseWriter, r *http.Request) {
	sessionID, _ := strconv.Atoi(chi.URLParam(r, "sessionID"))
	exercises, err := h.repo.GetExercises(r.Context(), sessionID)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if exercises == nil {
		exercises = []models.Exercise{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exercises)
}

// Create a new exercise in a session
func (h *ExerciseHandler) CreateExercise(w http.ResponseWriter, r *http.Request) {
	sessionID, _ := strconv.Atoi(chi.URLParam(r, "sessionID"))

	var ex models.Exercise
	if err := json.NewDecoder(r.Body).Decode(&ex); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	id, err := h.repo.CreateExercise(r.Context(), sessionID, ex)
	if err != nil {
		http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// Update exercise
func (h *ExerciseHandler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid exercise ID", http.StatusBadRequest)
		return
	}

	var req struct {
		Variation *string `json:"variation"`
		Notes     *string `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	ex := models.Exercise{
		Variation: req.Variation,
		Notes:     req.Notes,
	}

	if err := h.repo.UpdateExercise(r.Context(), id, ex); err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"id":     id,
		"status": "updated",
	})
}

// Delete exercise
func (h *ExerciseHandler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	if err := h.repo.DeleteExercise(r.Context(), id); err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{"status": "deleted", "id": id})
}
