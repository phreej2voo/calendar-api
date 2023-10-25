package response

import "calendar-api/models"

type Knowledge struct {
	Content     string `json:"content"`
	MoonAgeDesc string `json:"moonAgeDesc"`
	Date        string `json:"date"`
	Days        int    `json:"days"`
}

type Knowledges struct {
	Knowledges []Knowledge `json:"knowledges"`
	Paginate   Paginate    `json:"paginate"`
}

func NewKnowleges(baby models.Baby, calendarDays *[]models.CalendarDay, category string, paginate Paginate) *Response {
	knowleges := []Knowledge{}
	for _, calendarDay := range *calendarDays {
		content := calendarDay.ActionPoint
		if category == "viewpoint" {
			content = calendarDay.Viewpoint
		}
		knowleges = append(knowleges, Knowledge{
			Content:     content,
			MoonAgeDesc: calendarDay.MoonAgeDesc(),
			Date:        calendarDay.Date(baby),
			Days:        calendarDay.Day,
		})
	}
	return NewResponse(&Knowledges{
		Knowledges: knowleges,
		Paginate:   paginate,
	})
}
