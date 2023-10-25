package response

import "calendar-api/models"

type Login struct {
	JwtToken string `json:"jwtToken"`
	User
}

func NewLogin(user models.User) *Response {
	baby := user.CurrentBaby()

	return NewResponse(&Login{
		JwtToken: user.JwtToken(),
		User: User{
			NeedPhone: user.NeedPhone(),
			Baby: Baby{
				ID:       baby.ID,
				Name:     baby.Name,
				Birthday: baby.BirthdayStr(),
			},
		},
	})
}
