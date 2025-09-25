package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

// JSON marshaling (you already had this)
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))), nil
}

func (d *Date) UnmarshalJSON(b []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(b))
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

// --- Add these for pgx/sql compatibility ---
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	default:
		return fmt.Errorf("cannot scan %T into Date", value)
	}
}

func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

type Session struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	SessionDate time.Time  `json:"session_date"`
	SessionType string     `json:"session_type"`
	Notes       string     `json:"notes"`
	Exercises   []Exercise `json:"exercises,omitempty"`
}
