package controllers

import (
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func (CalendarDaysCtrl) Landing(c echo.Context) error {
	params := new(request.CalendarLanding)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	var query models.Query
	database.DB.Raw("SELECT MAX(day) as Max, MIN(day) as Min from calendar_days").Scan(&query)

	durationDays := int(time.Since(baby.Birthday).Hours()/24) + 1
	calenderDay := models.CalendarDay{}

	var result *gorm.DB
	if durationDays < query.Min {
		return NewCustomError(CalendarDayOutRange, "宝宝生日早于最小日历日期")
	} else if durationDays > query.Max {
		result = database.DB.Order("min_moon_age desc, day asc").Model(&models.CalendarDay{}).First(&calenderDay)
	} else {
		result = database.DB.Where(&models.CalendarDay{Day: durationDays}).First(&calenderDay)
	}

	if result.Error != nil {
		return NewCustomError(CalendarDayNotFound, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewCalendarDay(baby, calenderDay, true, params.AppVersion))
}

func (CalendarDaysCtrl) Show(c echo.Context) error {
	params := new(request.GetCalendarDay)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	durationDays := int(time.Until(params.Date.Time).Hours()/24) + 1
	if durationDays > 7 {
		return NewCustomError(CalendarDayLimit7Days, "最多可查看七天内的日历内容")
	}

	durationDays = int(params.Date.Sub(baby.Birthday).Hours()/24) + 1
	calenderDay := models.CalendarDay{}

	var query models.Query
	database.DB.Raw("SELECT MAX(day) as Max, MIN(day) as Min from calendar_days").Scan(&query)

	var result *gorm.DB
	if params.Date.Before(baby.Birthday) || durationDays < query.Min {
		return NewCustomError(CalendarDayOutRange, "查询日期早于最小日历日期")
	} else if durationDays > query.Max {
		return NewCustomError(CalendarDayOutRange, "查询日期超过最大日历日期")
	} else {
		result = database.DB.Where(&models.CalendarDay{Day: durationDays}).First(&calenderDay)
	}

	if result.Error != nil {
		return NewCustomError(CalendarDayNotFound, result.Error.Error())
	}

	return c.JSON(http.StatusOK, response.NewCalendarDay(baby, calenderDay, false, params.AppVersion))
}
