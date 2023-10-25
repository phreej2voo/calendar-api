package request

type SetPhoneNumber struct {
	Code string `json:"code" validate:"required"`
}
