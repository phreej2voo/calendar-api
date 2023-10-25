package response

import (
	"calendar-api/models"

	"github.com/labstack/echo/v4"
)

type Baby struct {
	ID          uint   `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Birthday    string `json:"birthday,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Phone       string `json:"phone,omitempty"`
	MoonAgeDesc string `json:"moonAgeDesc,omitempty"`
}

func NewBaby(baby models.Baby) *Response {
	return NewResponse(&Baby{
		ID:       baby.ID,
		Name:     baby.Name,
		Birthday: baby.BirthdayStr(),
	})
}

func ShowBaby(baby models.Baby, c echo.Context) *Response {
	user := c.Get("CurrentUser").(models.User)

	return NewResponse(&Baby{
		ID:       baby.ID,
		Name:     baby.Name,
		Birthday: baby.BirthdayStr(),
		Gender:   baby.Gender,
		Phone:    user.Phone,
	})
}
