package repository

import (
	"backend/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SetRepository struct {
	db *pgxpool.Pool
}

func NewSetRepository(db *pgxpool.Pool) *SetRepository {
	return &SetRepository{db: db}
}

// Get all sets for an exercise
func (r *SetRepository) GetSetsByExercise(ctx context.Context, exerciseID int) ([]models.Set, error) {
	rows, err := r.db.Query(ctx, `
		SELECT id, exercise_id, set_number, weight, weight_unit, duration, notes, reps
		FROM workout.sets
		WHERE exercise_id = $1
		ORDER BY set_number ASC
	`, exerciseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sets []models.Set
	for rows.Next() {
		var s models.Set
		var dur *time.Duration
		err := rows.Scan(
			&s.ID, &s.ExerciseID, &s.SetNumber,
			&s.Weight, &s.WeightUnit, &dur,
			&s.Notes, &s.Reps,
		)
		if dur != nil {
			ms := int64(*dur / time.Millisecond)
			s.Duration = &ms
		}
		if err != nil {
			return nil, err
		}
		sets = append(sets, s)
	}
	return sets, nil
}

// Create new set
func (r *SetRepository) CreateSet(ctx context.Context, s *models.Set) error {
	var dur *time.Duration
	if s.Duration != nil {
		d := time.Duration(*s.Duration) * time.Millisecond
		dur = &d
	}

	return r.db.QueryRow(ctx, `
    INSERT INTO workout.sets (exercise_id, set_number, weight, weight_unit, duration, notes, reps)
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    RETURNING id
`,
		s.ExerciseID, s.SetNumber, s.Weight, s.WeightUnit, dur, s.Notes, s.Reps,
	).Scan(&s.ID)
}

// Update existing set
func (r *SetRepository) UpdateSet(ctx context.Context, s *models.Set) error {
	var dur *time.Duration
	if s.Duration != nil {
		d := time.Duration(*s.Duration) * time.Millisecond
		dur = &d
	}
	_, err := r.db.Exec(ctx, `
		UPDATE workout.sets
		SET set_number=$1, weight=$2, weight_unit=$3, duration=$4, notes=$5, reps=$6
		WHERE id=$7
	`,
		s.SetNumber, s.Weight, s.WeightUnit, dur, s.Notes, s.Reps, s.ID,
	)
	return err
}

// Delete set
func (r *SetRepository) DeleteSet(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx, `DELETE FROM workout.sets WHERE id=$1`, id)
	return err
}
