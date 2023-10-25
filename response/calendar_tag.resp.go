package response

import (
	"calendar-api/database"
	"calendar-api/models"
)

type CalendarTag struct {
	ID           int       `json:"id,omitempty"`
	Name         string    `json:"name"`
	MoonAge      int       `json:"moonAge,omitempty"`
	AchievedID   int       `json:"achievedId,omitempty"`
	AchievedAt   string    `json:"achievedAt,omitempty"`
	AchievedDesc string    `json:"achievedDesc,omitempty"`
	Status       string    `json:"status,omitempty"`
	BgImg        string    `json:"bgImg,omitempty"`
	BgImgPin     string    `json:"bgImgPin,omitempty"`
	Contents     *[]string `json:"contents,omitempty"`
	BriefIntro   string    `json:"briefIntro,omitempty"`
	Baby         *Baby     `json:"baby,omitempty"`
}

type CalendarTags struct {
	Tags     []CalendarTag `json:"tags"`
	Moods    []models.Mood `json:"moods"`
	Paginate Paginate      `json:"paginate"`
}

func NewAchieveTags(babyTags []models.BabyCalendarTag, paginate Paginate) *Response {
	tags := []CalendarTag{}
	for _, babyTag := range babyTags {
		tags = append(tags, CalendarTag{
			ID:           int(babyTag.CalendarTag.ID),
			Name:         babyTag.CalendarTag.Name,
			AchievedID:   int(babyTag.ID),
			AchievedAt:   babyTag.CreatedAt.Format("2006-01-02"),
			AchievedDesc: babyTag.MoonAgeDesc(),
			BgImg:        babyTag.CalendarTag.Extras.BgImg,
			BgImgPin:     babyTag.CalendarTag.Extras.BgImgPin,
		})
	}

	return NewResponse(&CalendarTags{
		Tags:     tags,
		Paginate: paginate,
	})
}

func NewCalendarTags(calendarTags []models.CalendarTag, baby models.Baby, paginate Paginate) *Response {
	var calendarTagIds []uint
	for _, tag := range calendarTags {
		calendarTagIds = append(calendarTagIds, tag.ID)
	}

	var moods []models.Mood
	database.DB.Model(&models.Mood{}).Where("online = ?", true).Find(&moods)

	var babyCalendarTags []models.BabyCalendarTag
	database.DB.Model(&models.BabyCalendarTag{}).
		Where("baby_id = ? and calendar_tag_id in ?", baby.ID, calendarTagIds).
		Find(&babyCalendarTags)

	babyCalendarTagMap := make(map[int]models.BabyCalendarTag)
	for _, babyCalendarTag := range babyCalendarTags {
		babyCalendarTagMap[babyCalendarTag.CalendarTagID] = babyCalendarTag
	}

	babyMonth, _ := baby.MoonAge()

	tags := []CalendarTag{}
	for _, calendarTag := range calendarTags {
		var tagStatus string
		babyCalendarTag, ok := babyCalendarTagMap[int(calendarTag.ID)]
		if ok {
			tagStatus = "achieved"
		} else if calendarTag.StartMoonAge > babyMonth {
			tagStatus = "locked"
		} else {
			tagStatus = "unachieved"
		}
		tags = append(tags, CalendarTag{
			ID:         int(calendarTag.ID),
			Name:       calendarTag.Name,
			MoonAge:    calendarTag.StartMoonAge,
			BgImg:      calendarTag.Extras.BgImg,
			BgImgPin:   calendarTag.Extras.BgImgPin,
			Status:     tagStatus,
			AchievedID: int(babyCalendarTag.ID),
		})
	}
	return NewResponse(&CalendarTags{
		Tags:     tags,
		Paginate: paginate,
		Moods:    moods,
	})
}
