package request

type Login struct {
	Code       string `json:"code" validate:"required"`
	AchievedID int    `json:"achievedId"`
	UserID     int    `json:"userId"` // 注册邀请人
	Source     string `json:"source"` // 注册渠道
}
