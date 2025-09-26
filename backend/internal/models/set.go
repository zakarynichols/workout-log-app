package models

type Set struct {
	ID         int      `json:"id"`
	ExerciseID int      `json:"exercise_id"`
	SetNumber  int      `json:"set_number"`
	Weight     *float64 `json:"weight,omitempty"`
	WeightUnit string   `json:"weight_unit"`
	Reps       *int     `json:"reps,omitempty"`
	Duration   *int64   `json:"duration,omitempty"` // milliseconds
	Notes      *string  `json:"notes,omitempty"`
}
