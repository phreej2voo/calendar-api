package models

import "time"

type MonthTip struct {
	ID        uint      `json:"-"`
	MoonAge   int       `json:"-"`
	Name      string    `json:"name"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
