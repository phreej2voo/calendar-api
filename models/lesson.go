package models

import "time"

type Lesson struct {
	ID            uint      `json:"id"`
	MediaLink     string    `json:"mediaLink"`
	MediaType     string    `json:"mediaType"`
	MediaDuration int       `json:"mediaDuration"`
	LabelID       int       `json:"-"`
	StartDays     int       `json:"-"`
	EndDays       int       `json:"-"`
	CreatedAt     time.Time `json:"-"`
	UpdatedAt     time.Time `json:"-"`
}
