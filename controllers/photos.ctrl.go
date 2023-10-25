package controllers

import (
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (PhotosCtrl) Destroy(c echo.Context) error {
	user := c.Get("CurrentUser").(models.User)

	result := database.DB.Where("id = ? AND user_id = ?", c.Param("id"), user.ID).Delete(&models.Photo{})
	if result.RowsAffected == 0 {
		return NewCustomError(PhotoNotFound, "")
	}

	return c.JSON(http.StatusOK, response.NewResponse(nil))
}
