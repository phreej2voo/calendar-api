package controllers

import (
	"calendar-api/config"
	"calendar-api/database"
	"calendar-api/jobs"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"net/http"
	"time"

	"github.com/jrallison/go-workers"
	"github.com/labstack/echo/v4"
	"github.com/medivhzhan/weapp/v3/phonenumber"
)

func (UsersCtrl) SetPhoneNumber(c echo.Context) error {
	setPhoneReq := new(request.SetPhoneNumber)
	if err := BindValidate(c, setPhoneReq); err != nil {
		return err
	}

	user := c.Get("CurrentUser").(models.User)
	if len(user.Phone) > 0 {
		return NewCustomError(UserPhoneExistError, "")
	}

	phoneNumber := config.Weapp.Cli.NewPhonenumber()
	getPhoneRes, err := phoneNumber.GetPhoneNumber(&phonenumber.GetPhoneNumberRequest{Code: setPhoneReq.Code})

	if err != nil || getPhoneRes.CommonError.ErrCode != 0 {
		return NewCustomError(WXGetPhoneError, getPhoneRes.ErrMSG)
	}

	purePhoneNumber := getPhoneRes.Data.PurePhoneNumber
	countryCode := getPhoneRes.Data.CountryCode
	getPhoneAt := time.Unix(getPhoneRes.Data.Watermark.Timestamp, 0)

	result := database.DB.Model(&user).Updates(models.User{Phone: purePhoneNumber, GetPhoneAt: getPhoneAt, CountryCode: countryCode})
	if result.Error != nil {
		return NewCustomError(UpdateUserError, result.Error.Error())
	}

	workers.Enqueue(jobs.CrmLeadsQueue, "Add", purePhoneNumber)

	return c.JSON(http.StatusOK, response.NewResponse(nil))
}

func (UsersCtrl) Show(c echo.Context) error {
	user := c.Get("CurrentUser").(models.User)

	return c.JSON(http.StatusOK, response.NewUser(user))
}
