package controllers

import (
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (BabiesCtrl) Create(c echo.Context) error {
	params := new(request.CreateBaby)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	if time.Now().Before(params.Birthday.Time) {
		return NewCustomError(BabyBirthdayOutRange, "")
	}

	user := c.Get("CurrentUser").(models.User)
	baby := models.Baby{}
	database.DB.Model(&user).Association("Baby").Find(&baby)

	if baby.ID != 0 {
		return NewCustomError(BabyExistsError, "")
	}

	baby.Birthday = params.Birthday.Time
	baby.Name = params.Name
	baby.Gender = params.Gender
	baby.UserID = int(user.ID)

	result := database.DB.Create(&baby)
	if result.Error != nil {
		return NewCustomError(CreateBabyError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewBaby(baby))
}

func (BabiesCtrl) Show(c echo.Context) error {
	params := new(request.GetBaby)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response.ShowBaby(baby, c))
}

func (BabiesCtrl) Update(c echo.Context) error {
	params := new(request.UpdateBaby)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	user := c.Get("CurrentUser").(models.User)
	baby := models.Baby{}
	database.DB.Model(&user).Association("Baby").Find(&baby)
	if baby.ID == 0 {
		return NewCustomError(BabyNotExistsError, "")
	}

	if params.Name != "" {
		baby.Name = params.Name
	}
	if params.Gender != "" {
		baby.Gender = params.Gender
	}
	if !params.Birthday.IsZero() && params.Birthday.Time.Format("2006-01-02") != baby.Birthday.Format("2006-01-02") {
		baby.Birthday = params.Birthday.Time

		babyCalendarTags := models.BabyCalendarTag{}
		database.DB.Model(&babyCalendarTags).
			Where("baby_id = ?", baby.ID).
			Delete(&babyCalendarTags)
	}

	result := database.DB.Model(&baby).Updates(&baby)
	if result.Error != nil {
		return NewCustomError(UpdateBabyError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewResponse(nil))
}
