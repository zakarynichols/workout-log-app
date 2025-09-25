package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SessionRepository struct {
	db *pgxpool.Pool
}

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
	return &SessionRepository{db: db}
}

// --- DB Operations ---

func (r *SessionRepository) GetSessions(ctx context.Context, userID int) ([]models.Session, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, session_date, session_type, notes
        FROM workout.sessions
        WHERE user_id = $1
        ORDER BY session_date DESC, id DESC
    `, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []models.Session
	for rows.Next() {
		var s models.Session
		if err := rows.Scan(&s.ID, &s.SessionDate, &s.SessionType, &s.Notes); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	return sessions, nil
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID int, s models.Session) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, `
        INSERT INTO workout.sessions (user_id, session_date, session_type, notes)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, userID, s.SessionDate, s.SessionType, s.Notes).Scan(&id)
	return id, err
}

func (r *SessionRepository) UpdateSession(ctx context.Context, id int, s models.Session) error {
	_, err := r.db.Exec(ctx, `
        UPDATE workout.sessions
        SET session_date = $1, session_type = $2, notes = $3
        WHERE id = $4
    `, s.SessionDate, s.SessionType, s.Notes, id)
	return err
}

func (r *SessionRepository) DeleteSession(ctx context.Context, id int) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Delete sets
	_, err = tx.Exec(ctx, `
        DELETE FROM workout.sets
        WHERE exercise_id IN (SELECT id FROM workout.exercises WHERE session_id = $1)
    `, id)
	if err != nil {
		return err
	}

	// Delete exercises
	_, err = tx.Exec(ctx, `DELETE FROM workout.exercises WHERE session_id = $1`, id)
	if err != nil {
		return err
	}

	// Delete session
	_, err = tx.Exec(ctx, `DELETE FROM workout.sessions WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
