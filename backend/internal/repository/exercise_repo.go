package repository

import (
	"backend/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ExerciseRepository struct {
	db *pgxpool.Pool
}

func NewExerciseRepository(db *pgxpool.Pool) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

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

func (r *ExerciseRepository) CreateExercise(ctx context.Context, sessionID int, ex models.Exercise) (int, error) {
	var id int
	fmt.Printf("[Repo.CreateExercise] inserting sessionID=%d, dictID=%v, customID=%v, variation=%v, notes=%v\n",
		sessionID, ex.DictionaryExerciseID, ex.CustomExerciseID, ex.Variation, ex.Notes)

	err := r.db.QueryRow(ctx, `
        INSERT INTO workout.exercises (session_id, dictionary_exercise_id, custom_exercise_id, variation, notes)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id`,
		sessionID, ex.DictionaryExerciseID, ex.CustomExerciseID, ex.Variation, ex.Notes,
	).Scan(&id)

	if err != nil {
		fmt.Printf("[Repo.CreateExercise] error: %v\n", err)
		return 0, err
	}
	return id, nil
}

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

func (r *ExerciseRepository) DeleteExercise(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM workout.exercises WHERE id = $1`, id)
	return err
}

// Can this return the name?
func (r *ExerciseRepository) LookupOrCreateCustomExercise(ctx context.Context, userID int, name string) (int, error) {
	var id int
	err := r.db.QueryRow(ctx, `
        INSERT INTO workout.custom_exercises (user_id, name)
        VALUES ($1, $2)
        ON CONFLICT (user_id, name) DO UPDATE SET name = EXCLUDED.name
        RETURNING id
    `, userID, name).Scan(&id)
	return id, err
}

func (r *ExerciseRepository) GetDictionaryExerciseName(ctx context.Context, id int) (string, error) {
	var name string
	err := r.db.QueryRow(ctx, `SELECT name FROM workout.dictionary_exercises WHERE id = $1`, id).Scan(&name)
	return name, err
}

func (r *ExerciseRepository) GetCustomExerciseName(ctx context.Context, id int) (string, error) {
	var name string
	err := r.db.QueryRow(ctx, `SELECT name FROM workout.custom_exercises WHERE id = $1`, id).Scan(&name)
	return name, err
}
