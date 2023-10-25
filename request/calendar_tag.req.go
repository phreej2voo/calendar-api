package request

type CalendarTags struct {
	BabyID int `query:"babyId" validate:"required"`
	Page   int `query:"page"`
	Size   int `query:"size"`
}

func NewCalendarTags() *CalendarTags {
	return &CalendarTags{
		Page: 1,
		Size: 10,
	}
}
