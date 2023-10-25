package models

import "time"

type Photo struct {
	ID            int       `json:"id"`
	Link          string    `json:"link"`
	ImageableID   int       `json:"-"`
	ImageableType string    `json:"-"`
	UserID        int       `json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
