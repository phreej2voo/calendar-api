package models

import (
	"calendar-api/database"
	"calendar-api/tool"
	"fmt"
	"time"
)

type Baby struct {
	ID               uint
	Name             string
	Gender           string
	Birthday         time.Time `gorm:"type:date;default:null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	UserID           int
	CalendarTags     []CalendarTag `gorm:"many2many:baby_calendar_tags"`
	BabyCalendarTags []BabyCalendarTag
}

func (baby *Baby) BirthdayStr() string {
	if baby.Birthday.IsZero() {
		return ""
	}
	return baby.Birthday.Format("2006-01-02")
}

func (baby *Baby) MoonAge() (month, day int) {
	month, day = tool.MonthDuration(baby.Birthday, time.Now())
	return
}

func (baby *Baby) MoonAgeDesc() string {
	month, day := baby.MoonAge()
	return fmt.Sprintf("%d个月%d天", month, day)
}

func (baby *Baby) UnachievedTags(limit int) (tags *[]CalendarTag) {
	var achievedTagIds []uint
	database.DB.Model(&BabyCalendarTag{}).Where("baby_id = ?", int(baby.ID)).Pluck("calendar_tag_id", &achievedTagIds)
	month, _ := baby.MoonAge()
	database.DB.Order("start_moon_age desc").
		Model(&CalendarTag{}).
		Where("start_moon_age <= ?", month).
		Not(map[string]interface{}{"id": achievedTagIds}).
		Limit(limit).
		Find(&tags)
	return
}
