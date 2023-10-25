package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type Mood struct {
	ID        uint         `json:"id"`
	Name      string       `json:"name"`
	Contents  moodContents `json:"-" gorm:"type:json"`
	Online    bool         `json:"-"`
	CreatedAt time.Time    `json:"-"`
	UpdatedAt time.Time    `json:"-"`
}
type moodContents struct {
	AddPhoto []string `json:"add_photo"`
	Share    []string `json:"share"`
}

func (contents moodContents) Value() (driver.Value, error) {
	b, err := json.Marshal(contents)
	return string(b), err
}

func (contents *moodContents) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), contents)
}
