package request

type AchieveTag struct {
	CalendarTagID int `json:"tagId"`
	BabyID        int `json:"babyId"`
}

type CreateBabyCalendarTag struct {
	BabyID        int `json:"babyID" validate:"required"`
	CalendarTagID int `json:"tagId" validate:"required"`
	MoodID        int `json:"moodID"`
	AchievedAt    string
}
