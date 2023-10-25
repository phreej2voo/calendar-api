package models

import (
	"calendar-api/database"
	"calendar-api/tool"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type User struct {
	ID          uint `gorm:"primaryKey"`
	Phone       string
	GetPhoneAt  time.Time
	CountryCode string
	Openid      string `gorm:"unique" validate:"required"`
	Unionid     string
	AccessToken string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Baby        Baby
}

func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if len(user.AccessToken) == 0 {
		user.AccessToken, err = tool.GenerateRandomStringURLSafe(15)
	}
	return
}

func (user *User) NeedPhone() bool {
	return user.Phone == ""
}

func (user *User) JwtToken() (token string) {
	secret := os.Getenv("APP_SECRET")
	claims := jwt.MapClaims{}
	claims["user_id"] = user.ID
	claims["access_token"] = user.AccessToken
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, _ = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return
}

func (user *User) CurrentBaby() (baby Baby) {
	database.DB.Model(&user).Association("Baby").Find(&baby)
	return
}
