package response

import (
	"calendar-api/database"
	"calendar-api/models"
	"fmt"
	"math/rand"
	"time"
)

type CalendarDay struct {
	Today             string                   `json:"today,omitempty"`
	Days              int                      `json:"day"`
	Date              string                   `json:"date,omitempty"`
	Baby              Baby                     `json:"baby"`
	Better            string                   `json:"better"`
	Worse             string                   `json:"worse"`
	Viewpoint         string                   `json:"viewpoint"`
	ViewpointLabel    models.Label             `json:"viewpointLabel"`
	ViewpointLesson   models.Lesson            `json:"viewpointLesson"`
	ActionPoint       string                   `json:"actionPoint"`
	ActionPointLabel  models.Label             `json:"actionPointLabel"`
	ActionPointLesson models.Lesson            `json:"actionPointLesson"`
	MonthTip          models.MonthTip          `json:"title"` // 废弃：小程序二期上线后删除
	GrowthTitle       string                   `json:"growthTitle"`
	GrowthTargets     []models.GrowthTarget    `json:"growthTargets"` // 废弃： 小程序二期上线后删除，改用 growth indicators
	GrowthIndicators  []models.GrowthIndicator `json:"growthIndicators"`
	Tags              *[]CalendarTag           `json:"tags,omitempty"`
	DailySchedules    []models.DailySchedule   `json:"dailySchedules"`
}

func NewCalendarDay(baby models.Baby, calendarDay models.CalendarDay, landing bool, appVersion string) *Response {

	babyInfo := Baby{
		ID:          baby.ID,
		Name:        baby.Name,
		MoonAgeDesc: baby.MoonAgeDesc(),
		Birthday:    baby.BirthdayStr(),
	}

	betters := []string{"陪伴", "阳光浴", "对宝宝说话"}
	better, worse := calendarDay.Better, calendarDay.Worse
	if len(better) == 0 && len(worse) == 0 {
		better = betters[rand.Intn(len(betters))]
	}

	month, _ := baby.MoonAge()
	weekday := calendarDay.Weekday(baby)
	var dailySchedules []models.DailySchedule
	database.DB.Order("position asc").
		Where("min_moon_age <= ? and max_moon_age > ? and weekday = ?", month, month, weekday).
		Find(&dailySchedules)

	var viewpointLabel, actionPointLabel models.Label
	if calendarDay.ViewpointLabelID > 0 {
		database.DB.First(&viewpointLabel, calendarDay.ViewpointLabelID)
	}
	if calendarDay.ActionpointLabelID > 0 {
		database.DB.First(&actionPointLabel, calendarDay.ActionpointLabelID)
	}

	var viewpointLesson, actionPointLesson models.Lesson
	if calendarDay.ActionpointLabelID > 0 {
		database.DB.Where("start_days <= ? and end_days >= ? and label_id = ?", calendarDay.Day, calendarDay.Day, calendarDay.ActionpointLabelID).
			First(&actionPointLesson)
	}
	if calendarDay.ViewpointLabelID != calendarDay.ActionpointLabelID && calendarDay.ViewpointLabelID > 0 {
		database.DB.Where("start_days <= ? and end_days >= ? and label_id = ?", calendarDay.Day, calendarDay.Day, calendarDay.ViewpointLabelID).
			First(&viewpointLesson)
	}

	moonAge := calendarDay.Day / 30
	var growthIndicators []models.GrowthIndicator
	database.DB.Where("moon_age = ? and gender = ?", moonAge, baby.Gender).Find(&growthIndicators)

	// 二期小程序上线后，删除 growthTargets
	var growthTargets []models.GrowthTarget
	database.DB.Order("sequence asc").Model(&calendarDay).Association("GrowthTargets").Find(&growthTargets)

	var viewpoint string
	if appVersion == "v2" {
		viewpoint = calendarDay.Viewpoint
	} else if len(growthTargets) > 0 {
		viewpoint = ""
	}

	dayResponse := CalendarDay{
		Date:             calendarDay.Date(baby),
		Days:             calendarDay.Day,
		Baby:             babyInfo,
		Better:           better,
		Worse:            worse,
		Viewpoint:        viewpoint,
		ViewpointLabel:   viewpointLabel,
		ViewpointLesson:  viewpointLesson,
		ActionPointLabel: actionPointLabel,
		ActionPoint:      calendarDay.ActionPoint,
		GrowthTitle:      fmt.Sprintf("%d月龄身高体重标准", moonAge),
		GrowthTargets:    growthTargets,
		GrowthIndicators: growthIndicators,
		MonthTip:         calendarDay.MonthTip(),
		DailySchedules:   dailySchedules,
	}

	if landing {
		dayResponse.Today = time.Now().Format("2006-01-02")
		unachievedTags := *(baby.UnachievedTags(3))
		tags := []CalendarTag{}
		for _, unachievedTag := range unachievedTags {
			tags = append(tags, CalendarTag{
				ID:       int(unachievedTag.ID),
				Name:     unachievedTag.Name,
				BgImg:    unachievedTag.Extras.BgImg,
				BgImgPin: unachievedTag.Extras.BgImgPin,
			})
		}
		dayResponse.Tags = &tags
	}

	return NewResponse(&dayResponse)
}
