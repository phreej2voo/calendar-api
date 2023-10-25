package controllers

import (
	"calendar-api/database"
	"calendar-api/models"
	"calendar-api/request"
	"calendar-api/response"
	"calendar-api/tool"
	"fmt"
	"math"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (KnowledgeCtrl) Index(c echo.Context) error {
	params := request.NewKnowleges()
	if err := BindValidate(c, params); err != nil {
		return err
	}

	baby, err := CurrentBaby(c, uint(params.BabyID))
	if err != nil {
		return err
	}

	if !tool.SliceContains([]string{"viewpoint", "actionpoint"}, params.Category) {
		return NewCustomError(ParamError, "知识分类有误")
	}

	fields := map[string]string{"viewpoint": "viewpoint", "actionpoint": "action_point"}

	sql := fmt.Sprintf("%s_label_id = %d and %s is not null", params.Category, params.LabelID, fields[params.Category])

	currentPage := params.Page
	if params.Days >= 0 {
		var position int64
		database.DB.Model(&models.CalendarDay{}).Order("day asc").Where(sql).Where("day <= ?", params.Days).Count(&position)
		currentPage = int(math.Ceil(float64(position) / float64(params.Size)))
	}

	var calendarDays []models.CalendarDay
	database.DB.Order("day asc").Where(sql).Limit(params.Size).Offset((currentPage - 1) * params.Size).Find(&calendarDays)

	var totalCount int64
	database.DB.Model(&models.CalendarDay{}).Where(sql).Count(&totalCount)
	paginate := response.NewPaginate(currentPage, params.Size, int(totalCount))

	return c.JSON(http.StatusOK, response.NewKnowleges(baby, &calendarDays, params.Category, paginate))
}
