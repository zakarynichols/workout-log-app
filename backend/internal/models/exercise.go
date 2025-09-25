package models

type Exercise struct {
	ID                   int     `json:"-"`
	Name                 string  `json:"name"`
	DictionaryExerciseID *int    `json:"-"`
	CustomExerciseID     *int    `json:"-"`
	Variation            *string `json:"variation,omitempty"`
	Notes                *string `json:"notes,omitempty"`
	Sets                 []Set   `json:"sets"`
}
