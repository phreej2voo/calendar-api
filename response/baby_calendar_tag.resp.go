package response

import (
	"calendar-api/database"
	"calendar-api/models"
	"fmt"
	"math/rand"

	"github.com/labstack/echo/v4"
)

type CreateBabyCalendarTag struct {
	TagID        int           `json:"tagId"`
	Name         string        `json:"name"`
	MoonAge      int           `json:"moonAge"`
	Status       string        `json:"status"`
	BgImg        string        `json:"bgImg"`
	BgImgPin     string        `json:"bgImgPin"`
	AchievedAt   string        `json:"achievedAt"`
	AchievedDesc string        `json:"achievedDesc"`
	AchievedID   int           `json:"achievedId"`
	Tags         []CalendarTag `json:"tags"`
}

func NewCreateBabyCalendarTag(babyTag models.BabyCalendarTag) *Response {
	var baby models.Baby
	database.DB.Model(&babyTag).Association("Baby").Find(&baby)
	calendarTags := baby.UnachievedTags(3)
	var tags []CalendarTag
	for _, calendarTag := range *calendarTags {
		tags = append(tags, CalendarTag{
			ID:       int(calendarTag.ID),
			Name:     calendarTag.Name,
			BgImg:    calendarTag.Extras.BgImg,
			BgImgPin: calendarTag.Extras.BgImgPin,
		})
	}

	var calendarTag models.CalendarTag
	database.DB.Model(&babyTag).Association("CalendarTag").Find(&calendarTag)

	return NewResponse(&CreateBabyCalendarTag{
		TagID:        int(calendarTag.ID),
		Name:         calendarTag.Name,
		MoonAge:      calendarTag.StartMoonAge,
		Status:       "achieved",
		BgImg:        calendarTag.Extras.BgImg,
		BgImgPin:     calendarTag.Extras.BgImgPin,
		AchievedAt:   babyTag.AchievedAt.Format("2006-01-02"),
		AchievedDesc: babyTag.MoonAgeDesc(),
		AchievedID:   int(babyTag.ID),
		Tags:         tags,
	})
}

type ShowBabyCalendarTag struct {
	IsSelf       bool           `json:"isSelf"`
	Baby         Baby           `json:"baby"`
	Tag          CalendarTag    `json:"tag"`
	Mood         models.Mood    `json:"mood"`
	Photos       []models.Photo `json:"photos"`
	AddPhotoText string         `json:"addPhotoText"`
}

func NewShowBabyCalendarTag(babyCalendarTag models.BabyCalendarTag, c echo.Context) *Response {
	user := c.Get("CurrentUser").(models.User)
	month, day := babyCalendarTag.MoonAge()

	addPhotoText := "当然要拍照记录呀，ta又不是永远无忧无虑"
	if addPhotoLen := len(babyCalendarTag.Mood.Contents.AddPhoto); addPhotoLen > 0 {
		addPhotoText = babyCalendarTag.Mood.Contents.AddPhoto[rand.Intn(addPhotoLen)]
	}

	return NewResponse(&ShowBabyCalendarTag{
		IsSelf: user.ID == uint(babyCalendarTag.Baby.UserID),
		Baby: Baby{
			ID:          babyCalendarTag.Baby.ID,
			Name:        babyCalendarTag.Baby.Name,
			MoonAgeDesc: fmt.Sprintf("%d月龄%d天", month, day),
		},
		Tag: CalendarTag{
			Name:       babyCalendarTag.CalendarTag.Name,
			BriefIntro: babyCalendarTag.CalendarTag.BriefIntro,
			Contents:   &babyCalendarTag.CalendarTag.Extras.Contents,
		},
		Mood: models.Mood{
			ID:   babyCalendarTag.Mood.ID,
			Name: babyCalendarTag.Mood.Name,
		},
		Photos:       babyCalendarTag.Photos,
		AddPhotoText: addPhotoText,
	})
}

type ShareBabyCalendarTag struct {
	QrCodeBase64 string         `json:"qrCodeBase64"`
	ShareText    string         `json:"shareText"`
	ShareTexts   []string       `json:"shareTexts"`
	Photos       []models.Photo `json:"photos"`
	BriefIntro   string         `json:"briefInto"`
}

func NewShareBabyCalendarTag(babyCalendarTag models.BabyCalendarTag, qrCodeBase64 string, c echo.Context) *Response {
	mood := babyCalendarTag.Mood

	shareTexts := []string{"可爱如你，幽默如你，搞怪如你", "生活明朗 万物可爱 按时长大"}
	if shareLen := len(mood.Contents.Share); shareLen > 0 {
		shareTexts = mood.Contents.Share
	}
	shareText := shareTexts[rand.Intn(len(shareTexts))]

	return NewResponse(&ShareBabyCalendarTag{
		QrCodeBase64: qrCodeBase64,
		ShareText:    shareText,
		ShareTexts:   shareTexts,
		BriefIntro:   babyCalendarTag.CalendarTag.BriefIntro,
		Photos:       babyCalendarTag.Photos,
	})
}
