package models

import (
	"calendar-api/database"
	"calendar-api/tool"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type BabyCalendarTag struct {
	ID            uint
	BabyID        int `json:"babyID" validate:"required"`
	CalendarTagID int `json:"tagId" validate:"required"`
	MoodID        int `json:"moodID"`
	AchievedAt    time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
	CalendarTag   CalendarTag `gorm:"foreignKey:CalendarTagID;references:ID"`
	Baby          Baby
	Photos        []Photo `gorm:"polymorphic:Imageable;polymorphicValue:BabyCalendarTag"`
	Mood          Mood
}

func (babyTag *BabyCalendarTag) MoonAge() (month, day int) {
	if babyTag.Baby.ID == 0 {
		database.DB.Model(babyTag).Association("Baby").Find(&babyTag.Baby)
	}

	month, day = tool.MonthDuration(babyTag.Baby.Birthday, babyTag.CreatedAt)
	return
}

func (babyTag *BabyCalendarTag) MoonAgeDesc() string {
	if babyTag.Baby.ID == 0 {
		database.DB.Model(babyTag).Association("Baby").Find(&babyTag.Baby)
	}
	month, day := babyTag.MoonAge()
	return fmt.Sprintf("%d月龄%d天 %s获得了", month, day, babyTag.Baby.Name)
}
