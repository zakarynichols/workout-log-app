package models

type Exercise struct {
	ID                   int     `json:"id"`
	SessionID            int     `json:"session_id"`
	DictionaryExerciseID *int    `json:"dictionary_exercise_id,omitempty"`
	CustomExerciseID     *int    `json:"custom_exercise_id,omitempty"`
	Variation            *string `json:"variation,omitempty"`
	Notes                *string `json:"notes,omitempty"`
	Name                 *string `json:"name,omitempty"`
}
