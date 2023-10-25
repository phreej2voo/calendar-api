package controllers

import (
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 所有标签列表
func (CalendarTagsCtrl) Index(c echo.Context) error {
	params := request.NewCalendarTags()
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	tags := []models.CalendarTag{}
	database.DB.Order("start_moon_age asc, id asc").
		Limit(params.Size).
		Offset((params.Page - 1) * params.Size).
		Find(&tags)

	var tagsCount int64
	database.DB.Model(&models.CalendarTag{}).Count(&tagsCount)
	paginate := response.NewPaginate(params.Page, params.Size, int(tagsCount))

	return c.JSON(http.StatusOK, response.NewCalendarTags(tags, baby, paginate))
}

// 已领取标签列表
func (CalendarTagsCtrl) Achieved(c echo.Context) error {
	params := request.NewCalendarTags()
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	var babyCalendarTags []models.BabyCalendarTag
	database.DB.Preload("CalendarTag").Preload("Baby").Order("id desc").
		Limit(params.Size).
		Offset((params.Page - 1) * params.Size).
		Model(&baby).
		Association("BabyCalendarTags").
		Find(&babyCalendarTags)

	var babyTagsCount int64
	database.DB.Model(&models.BabyCalendarTag{}).Where("baby_id = ?", baby.ID).Count(&babyTagsCount)
	paginate := response.NewPaginate(params.Page, params.Size, int(babyTagsCount))

	return c.JSON(http.StatusOK, response.NewAchieveTags(babyCalendarTags, paginate))
}
