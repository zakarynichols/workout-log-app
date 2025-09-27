package handlers

import (
	"backend/internal/models"
	"backend/internal/repository"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type SetHandler struct {
	repo *repository.SetRepository
}

func NewSetHandler(repo *repository.SetRepository) *SetHandler {
	return &SetHandler{repo: repo}
}

func (h *SetHandler) GetSets(w http.ResponseWriter, r *http.Request) {
	exerciseID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	sets, err := h.repo.GetSetsByExercise(r.Context(), exerciseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(sets)
}

func (h *SetHandler) CreateSet(w http.ResponseWriter, r *http.Request) {
	exerciseID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s models.Set
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.ExerciseID = exerciseID
	if err := h.repo.CreateSet(r.Context(), &s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func (h *SetHandler) UpdateSet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s models.Set
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.ID = id
	if err := h.repo.UpdateSet(r.Context(), &s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"id": id, "status": "updated"})
}

func (h *SetHandler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.repo.DeleteSet(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]any{"id": id, "status": "deleted"})
}
