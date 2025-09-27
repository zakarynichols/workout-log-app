package handlers

import (
	"backend/internal/models"
	"backend/internal/repository"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	// Format as "YYYY-MM-DD"
	formatted := fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}

// getUserID: single place to swap logic later
func getUserID(r *http.Request) int {
	return 1 // DEV ONLY: Hardcoded until auth is added
}

type SessionHandler struct {
	repo *repository.SessionRepository
}

func NewSessionHandler(repo *repository.SessionRepository) *SessionHandler {
	return &SessionHandler{repo: repo}
}

func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	session, err := h.repo.GetSession(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Session not found", http.StatusNotFound)
		} else {
			http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (h *SessionHandler) GetSessions(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)

	sessions, err := h.repo.GetSessions(r.Context(), userID)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// Request type for both POST and PUT session
type SessionRequest struct {
	SessionDate models.Date `json:"session_date"`
	SessionType string      `json:"session_type"`
	Notes       string      `json:"notes"`
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request) {
	userID := getUserID(r)

	var req SessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	session := models.Session{
		UserID:      userID,
		SessionDate: req.SessionDate.Time,
		SessionType: req.SessionType,
		Notes:       req.Notes,
	}

	id, err := h.repo.CreateSession(r.Context(), userID, session)
	if err != nil {
		http.Error(w, "DB insert error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *SessionHandler) UpdateSession(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var req SessionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	s := models.Session{
		ID:          id,
		SessionDate: req.SessionDate.Time,
		SessionType: req.SessionType,
		Notes:       req.Notes,
	}

	updated, err := h.repo.UpdateSession(r.Context(), id, s)
	if err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(updated); err != nil {
		http.Error(w, "Encoding error: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	id, _ := strconv.Atoi(idStr)

	if err := h.repo.DeleteSession(r.Context(), id); err != nil {
		http.Error(w, "DB error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
