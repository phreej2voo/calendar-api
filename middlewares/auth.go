package middlewares

import (
	"calendar-api/database"
	"calendar-api/models"
	"fmt"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var KeyAuthConfig = middleware.KeyAuthConfig{
	KeyLookup: "header:Authorization",
	Validator: func(key string, c echo.Context) (bool, error) {
		tokenStr := strings.Split(c.Request().Header.Get("Authorization"), " ")[1]

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, err error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil {
			return false, err
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !(ok && token.Valid) {
			return false, nil
		}

		user := models.User{}
		database.DB.
			Where(&models.User{ID: uint(claims["user_id"].(float64)), AccessToken: claims["access_token"].(string)}).
			First(&user)

		if user.ID == 0 {
			return false, nil
		}
		c.Set("CurrentUser", user)
		return true, nil
	},
	Skipper: func(c echo.Context) bool {
		return c.Request().RequestURI == "/app/api/login"
	},
}
