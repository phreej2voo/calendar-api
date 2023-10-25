package response

import "calendar-api/models"

type User struct {
	NeedPhone bool `json:"needPhone"`
	Baby      Baby `json:"baby"`
}

func NewUser(user models.User) *Response {
	baby := user.CurrentBaby()

	return NewResponse(&User{
		NeedPhone: user.NeedPhone(),
		Baby: Baby{
			ID:       baby.ID,
			Name:     baby.Name,
			Birthday: baby.BirthdayStr(),
		},
	})
}
