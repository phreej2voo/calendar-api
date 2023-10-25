package middlewares

import (
	"bytes"
	"calendar-api/tool"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type responseBodyDump struct {
	io.Writer
	http.ResponseWriter
}

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		reqBody := []byte{}
		if c.Request().Body != nil {
			reqBody, _ = ioutil.ReadAll(c.Request().Body)
		}
		c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		reqHeader := c.Request().Header

		resBody := new(bytes.Buffer)
		mw := io.MultiWriter(c.Response().Writer, resBody)
		writer := &responseBodyDump{Writer: mw, ResponseWriter: c.Response().Writer}
		c.Response().Writer = writer

		start := time.Now()

		if err = next(c); err != nil {
			c.Error(err)
		}

		requestId := reqHeader.Get(echo.HeaderXRequestID)
		if requestId == "" {
			requestId = c.Response().Header().Get(echo.HeaderXRequestID)
		}

		tool.Logger.WithFields(logrus.Fields{
			"uri":           c.Request().RequestURI,
			"method":        c.Request().Method,
			"User-Agent":    c.Request().UserAgent(),
			"user_id":       userID(reqHeader),
			"protocol":      c.Request().Proto,
			"request":       string(reqBody),
			"request_id":    requestId,
			"remote_ip":     c.RealIP(),
			"latency_human": time.Since(start).Nanoseconds(),
			"status":        c.Response().Status,
			"response":      resBody.String(),
		}).Info("CALENDER-API")
		return
	}
}

func (w *responseBodyDump) WriteHeader(code int) {
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseBodyDump) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func userID(reqHeader http.Header) float64 {
	tokens := strings.Split(reqHeader.Get("Authorization"), " ")
	tokenString := tokens[len(tokens)-1]
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("APP_SECRET")), nil
	})
	if claims["user_id"] != nil {
		return claims["user_id"].(float64)
	}
	return 0
}
