package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

func (ce *CustomError) Error() string {
	return ce.Message
}

const (
	ParamError int = 10101

	CreateUserError     int = 20101
	UpdateUserError     int = 20102
	UserPhoneExistError int = 20103

	CreateBabyError      int = 20201
	BabyNotFound         int = 20202
	BabyExistsError      int = 20203
	BabyBirthdayOutRange int = 20204
	UpdateBabyError      int = 20205
	BabyNotExistsError   int = 20206

	CalendarDayNotFound   int = 20301
	CalendarDayOutRange   int = 20302
	CalendarDayLimit7Days int = 20303

	CreateBabyCalendarTagError int = 20401
	BabyCalendarTagNotFound    int = 20402

	CalendarTagNotFound int = 20501

	PhotoNotFound    int = 20550
	CreatePhotoError int = 205501

	WeappError          int = 40101
	WXCode2SessionError int = 40102
	WXGetPhoneError     int = 40103
	WxGetQRCodeError    int = 40104
)

var codeMessage = map[int]string{
	ParamError:                 "参数信息有误",
	WXCode2SessionError:        "微信登录凭证校验失败",
	WXGetPhoneError:            "微信获取手机号失败",
	WxGetQRCodeError:           "小程序码获取失败",
	CreateUserError:            "用户信息创建失败",
	UpdateUserError:            "用户信息更新失败",
	CreateBabyError:            "宝宝信息创建失败",
	UpdateBabyError:            "宝宝信息更新失败",
	BabyNotFound:               "查询宝宝信息失败",
	BabyBirthdayOutRange:       "本产品暂未提供预产期的宝宝相关内容",
	CalendarDayNotFound:        "日历日查询失败",
	CreateBabyCalendarTagError: "标签领取失败",
	BabyCalendarTagNotFound:    "标签领取记录查询失败",
	UserPhoneExistError:        "已获取过用户手机号",
	BabyExistsError:            "宝宝已创建",
	BabyNotExistsError:         "宝宝未创建",
	CalendarDayOutRange:        "日期超出范围",
	CalendarTagNotFound:        "日历标签查询失败",
	CalendarDayLimit7Days:      "最多往后看7天内的内容",
	PhotoNotFound:              "照片信息查找失败",
}

func NewCustomError(code int, detail string) error {
	return &CustomError{
		Code:    code,
		Message: codeMessage[code],
		Detail:  detail,
	}
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if he, ok := err.(*CustomError); ok {
		c.JSON(http.StatusOK, he)
	} else {
		c.Echo().DefaultHTTPErrorHandler(err, c)
	}
}
