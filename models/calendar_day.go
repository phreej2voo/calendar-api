package models

import (
	"calendar-api/database"
	"fmt"
	"time"
)

type CalendarDay struct {
	ID                 uint
	MinMoonAge         int
	MaxMoonAge         int
	Day                int
	Better             string
	Worse              string
	Viewpoint          string
	ActionPoint        string
	CreatedAt          time.Time
	UpdatedAt          time.Time
	GrowthTargets      []GrowthTarget
	ViewpointLabelID   int
	ActionpointLabelID int
}

func (calendar *CalendarDay) MonthTip() (monthTip MonthTip) {
	database.DB.Model(&monthTip).Where("moon_age = ?", calendar.MinMoonAge).First(&monthTip)
	return
}

func (calendar *CalendarDay) Date(baby Baby) (date string) {
	date = baby.Birthday.AddDate(0, 0, calendar.Day-1).Format("2006-01-02")
	return
}

func (calendar *CalendarDay) MoonAgeDesc() string {
	month, day := calendar.Day/30, calendar.Day%30
	return fmt.Sprintf("%d个月%d天", month, day)
}

func (calendar *CalendarDay) Weekday(baby Baby) (weekday string) {
	weekday = baby.Birthday.AddDate(0, 0, calendar.Day-1).Weekday().String()
	return
}
