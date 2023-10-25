package tool

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Validation struct {
	Validator *validator.Validate
}

func (v *Validation) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}

func NewValidation() echo.Validator {
	return &Validation{Validator: validator.New()}
}
