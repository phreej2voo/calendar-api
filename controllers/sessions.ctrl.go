package controllers

import (
	"calendar-api/config"
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/medivhzhan/weapp/v3/auth"
	"gorm.io/gorm"
)

// 微信小程序登陆
func (SessionsCtrl) Create(c echo.Context) error {
	params := new(request.Login)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	res, err := config.Weapp.Auth.Code2Session(&auth.Code2SessionRequest{
		Appid:     os.Getenv("MP_APPID"),
		Secret:    os.Getenv("MP_SECRET"),
		JsCode:    params.Code,
		GrantType: "authorization_code",
	})

	if err != nil || res.CommonError.ErrCode != 0 {
		return NewCustomError(WXCode2SessionError, res.CommonError.ErrMSG)
	}

	user := models.User{}
	result := database.DB.Where(models.User{Openid: res.Openid}).Assign(models.User{Unionid: res.Unionid}).FirstOrCreate(&user)
	if result.Error != nil {
		return NewCustomError(CreateUserError, result.Error.Error())
	}

	if params.AchievedID != 0 {
		var babyCalendarTag models.BabyCalendarTag
		database.DB.Preload("Baby").First(&babyCalendarTag, params.AchievedID)

		if babyCalendarTag.ID != 0 && babyCalendarTag.Baby.UserID != int(user.ID) {
			var invitation models.Invitation
			database.DB.FirstOrCreate(&invitation, models.Invitation{
				InviterID:     babyCalendarTag.Baby.UserID,
				InviteeID:     int(user.ID),
				InvitableID:   int(babyCalendarTag.ID),
				InvitableType: "BabyCalendarTag",
				Source:        params.Source,
			})
		}
	} else if params.UserID != 0 {
		var invitation models.Invitation
		result := database.DB.Where("inviter_id = ? AND invitee_id = ?", params.UserID, user.ID).First(&invitation)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			database.DB.Create(&models.Invitation{
				InviterID: params.UserID,
				InviteeID: int(user.ID),
				Source:    params.Source,
			})
		}
	}

	return c.JSON(http.StatusOK, response.NewLogin(user))
}
