package repository

import (
	"backend/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExerciseRepository struct {
	db *pgxpool.Pool
}

func NewExerciseRepository(db *pgxpool.Pool) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

// Get all exercises for a session
func (r *ExerciseRepository) GetExercises(ctx context.Context, sessionID int) ([]models.Exercise, error) {
	rows, err := r.db.Query(ctx, `
        SELECT id, session_id, dictionary_exercise_id, custom_exercise_id, variation, notes
        FROM workout.exercises
        WHERE session_id = $1
        ORDER BY id`, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var exercises []models.Exercise
	for rows.Next() {
		var ex models.Exercise
		err := rows.Scan(&ex.ID, &ex.SessionID, &ex.DictionaryExerciseID, &ex.CustomExerciseID, &ex.Variation, &ex.Notes)
		if err != nil {
			return nil, err
		}
		exercises = append(exercises, ex)
	}
	return exercises, rows.Err()
}

// Create exercise
func (r *ExerciseRepository) CreateExercise(ctx context.Context, sessionID int, ex models.Exercise) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, `
        INSERT INTO workout.exercises (session_id, dictionary_exercise_id, custom_exercise_id, variation, notes)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`,
		sessionID, ex.DictionaryExerciseID, ex.CustomExerciseID, ex.Variation, ex.Notes,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Update exercise
func (r *ExerciseRepository) UpdateExercise(ctx context.Context, id int, e models.Exercise) error {
	_, err := r.db.Exec(ctx, `
        UPDATE workout.exercises
        SET
            variation = COALESCE($1, variation),
            notes = COALESCE($2, notes)
        WHERE id = $3
    `, e.Variation, e.Notes, id)

	return err
}

// Delete exercise
func (r *ExerciseRepository) DeleteExercise(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM workout.exercises WHERE id = $1`, id)
	return err
}
