package models

import "time"

type GrowthIndicator struct {
	ID        int       `json:"-"`
	MoonAge   int       `json:"-"`
	Gender    string    `json:"-"`
	Dimension string    `json:"dimension"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
