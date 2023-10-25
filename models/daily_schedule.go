package models

import "time"

type DailySchedule struct {
	ID         uint      `json:"-"`
	MinMoonAge int       `json:"-"`
	MaxMoonAge int       `json:"-"`
	Weekday    string    `json:"-"`
	Hour       string    `json:"hour"`
	Item       string    `json:"item"`
	Content    string    `json:"content"`
	Position   int       `json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
}
