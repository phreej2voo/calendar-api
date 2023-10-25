package controllers

import (
	"calendar-api/database"
	"calendar-api/models"

	"github.com/labstack/echo/v4"
)

func CurrentBaby(c echo.Context, babyID uint) (baby models.Baby, err error) {
	user := c.Get("CurrentUser").(models.User)
	result := database.DB.Where(&models.Baby{UserID: int(user.ID), ID: babyID}).First(&baby)
	if baby.ID == 0 {
		err = NewCustomError(BabyNotFound, result.Error.Error())
	}
	return
}
