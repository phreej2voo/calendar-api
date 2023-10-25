package request

type GetCalendarDay struct {
	BabyID     int        `query:"babyId" validate:"required"`
	Date       CustomDate `query:"date" validate:"required"`
	AppVersion string     `query:"appVersion"`
}

type CalendarLanding struct {
	BabyID     int    `query:"babyId" validate:"required"`
	AppVersion string `query:"appVersion"`
}
