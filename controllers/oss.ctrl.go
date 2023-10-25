package controllers

import (
	"calendar-api/request"
	"calendar-api/response"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
)

var expireTime int64 = 3600

type Policy struct {
	Expiration string        `json:"expiration"`
	Conditions []interface{} `json:"conditions"`
}

func (OssCtrl) getAliyunPolicy(c echo.Context) error {
	params := new(request.OssAliyunPolicy)
	if err := BindValidate(c, params); err != nil {
		return err
	}

	var policy Policy

	dir := os.Getenv("ALIYUN_UPLOAD_DIR")
	key := fmt.Sprintf(dir, params.ModelName, params.ModelId)

	condition := []string{"starts-with", "$key", key}
	policy.Conditions = append(policy.Conditions, map[string]string{"bucket": os.Getenv("ALIYUN_BUCKET")})
	policy.Conditions = append(policy.Conditions, condition)
	policy.Expiration = time.Unix(time.Now().Unix()+expireTime, 0).UTC().Format("2006-01-02T15:04:05Z")

	result, _ := json.Marshal(policy)
	encodedPolicy := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(os.Getenv("ALIYUN_ACCESS_KEY")))
	io.WriteString(h, encodedPolicy)
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return c.JSON(http.StatusOK, response.NewAliyunOssPolicy(key, encodedPolicy, signature))
}
