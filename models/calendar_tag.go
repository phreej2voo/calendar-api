package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type CalendarTag struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	StartMoonAge int       `json:"-"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Extras       tagExtras `json:"-" gorm:"type:json"`
	BriefIntro   string    `json:"briefIntro"`
}

type tagExtras struct {
	BgImg    string   `json:"bg_img"`
	BgImgPin string   `json:"bg_img_pin"`
	Contents []string `json:"contents"`
}

func (extras tagExtras) Value() (driver.Value, error) {
	b, err := json.Marshal(extras)
	return string(b), err
}

func (extras *tagExtras) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), extras)
}
