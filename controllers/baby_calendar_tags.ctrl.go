package controllers

import (
	"calendar-api/config"
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/medivhzhan/weapp/v3"
)

func (BabyCalendarTagCtrl) Create(c echo.Context) error {
	CreateBabyCalendarTag := new(request.CreateBabyCalendarTag)
	if err := BindValidate(c, CreateBabyCalendarTag); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(CreateBabyCalendarTag.BabyID))
	if baby.ID == 0 {
		return err
	}

	var achievedAt time.Time
	if achievedAt, err = time.Parse("2006-01-02", CreateBabyCalendarTag.AchievedAt); err != nil {
		achievedAt = time.Now()
	}
	if achievedAt.Before(baby.Birthday) || achievedAt.After(time.Now()) {
		return NewCustomError(ParamError, "日期错误")
	}

	babyCalendarTag := new(models.BabyCalendarTag)
	babyCalendarTag.AchievedAt = achievedAt
	babyCalendarTag.BabyID = CreateBabyCalendarTag.BabyID
	babyCalendarTag.MoodID = CreateBabyCalendarTag.MoodID
	babyCalendarTag.CalendarTagID = CreateBabyCalendarTag.CalendarTagID

	result := database.DB.Where(babyCalendarTag).FirstOrCreate(&babyCalendarTag)
	if result.Error != nil {
		return NewCustomError(CreateBabyCalendarTagError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewCreateBabyCalendarTag(*babyCalendarTag))
}

func (BabyCalendarTagCtrl) Show(c echo.Context) error {
	var babyCalendarTag models.BabyCalendarTag
	result := database.DB.Preload("Baby").Preload("CalendarTag").Preload("Mood").Preload("Photos").First(&babyCalendarTag, c.Param("id"))
	if babyCalendarTag.ID == 0 {
		return NewCustomError(BabyCalendarTagNotFound, result.Error.Error())
	}
	if babyCalendarTag.Baby.ID == 0 {
		return NewCustomError(BabyNotFound, result.Error.Error())
	}
	if babyCalendarTag.CalendarTag.ID == 0 {
		return NewCustomError(CalendarTagNotFound, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewShowBabyCalendarTag(babyCalendarTag, c))
}

func (BabyCalendarTagCtrl) AddPhoto(c echo.Context) error {
	user := c.Get("CurrentUser").(models.User)

	var babyCalendarTag models.BabyCalendarTag
	result := database.DB.Joins("left join babies on babies.id = baby_calendar_tags.baby_id").
		Where("baby_calendar_tags.id = ? and babies.user_id = ?", c.Param("id"), user.ID).
		First(&babyCalendarTag)
	if babyCalendarTag.ID == 0 {
		return NewCustomError(BabyCalendarTagNotFound, result.Error.Error())
	}

	params := new(request.AddPhotos)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	photos := []models.Photo{}
	for _, link := range params.Links {
		photos = append(photos, models.Photo{
			ImageableID:   int(babyCalendarTag.ID),
			ImageableType: "BabyCalendarTag",
			UserID:        int(user.ID),
			Link:          link,
		})
	}

	if len(photos) > 0 {
		database.DB.Create(&photos)
	}

	return c.JSON(http.StatusOK, response.NewResponse(&photos))
}

func (BabyCalendarTagCtrl) Share(c echo.Context) error {
	var babyCalendarTag models.BabyCalendarTag
	result := database.DB.Preload("Baby").Preload("CalendarTag").Preload("Mood").Preload("Photos").First(&babyCalendarTag, c.Param("id"))
	if babyCalendarTag.ID == 0 {
		return NewCustomError(BabyCalendarTagNotFound, result.Error.Error())
	}

	resp, ce, err := config.Weapp.Cli.GetUnlimitedQRCode(&weapp.UnlimitedQRCode{
		Scene: fmt.Sprintf("achievedId=%d", babyCalendarTag.ID),
		Page:  "pages/tag/detail",
	})

	if err != nil || ce.ErrCode != 0 {
		return NewCustomError(WxGetQRCodeError, ce.ErrMSG)
	}

	readBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return NewCustomError(WxGetQRCodeError, ce.ErrMSG)
	}
	resp.Body.Close()

	qrCodeBase64 := base64.StdEncoding.EncodeToString(readBody)

	return c.JSON(http.StatusOK, response.NewShareBabyCalendarTag(babyCalendarTag, qrCodeBase64, c))
}
