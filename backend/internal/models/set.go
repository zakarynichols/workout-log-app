package models

type Set struct {
	SetNumber  int      `json:"set_number"`
	Weight     *float64 `json:"weight,omitempty"`
	WeightUnit string   `json:"weight_unit,omitempty"`
	Reps       *int     `json:"reps,omitempty"`
	DurationMS int64    `json:"duration_ms,omitempty"`
	Notes      string   `json:"notes,omitempty"`
}
