package response

import (
	"os"
)

type AliyunOssPolicy struct {
	Key         string `json:"key"`
	Url         string `json:"url"`
	Policy      string `json:"policy"`
	Signature   string `json:"signature"`
	AccessKeyID string `json:"accessKeyId"`
}

func NewAliyunOssPolicy(key, policy, signature string) *Response {

	return NewResponse(&AliyunOssPolicy{
		Key:         key,
		Url:         os.Getenv("ALIYUN_UPLOAD_HOST"),
		Policy:      policy,
		Signature:   signature,
		AccessKeyID: os.Getenv("ALIYUN_ACCESS_ID"),
	})
}
