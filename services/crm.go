package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

func CrmGet(url string, queryValues url.Values) map[string]interface{} {
	req, _ := http.NewRequest("GET", os.Getenv("CRM_HOST")+url, nil)

	req.Header.Add("Authorization", os.Getenv("CRM_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	req.URL.RawQuery = queryValues.Encode()

	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	var respBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)
	return respBody
}

func CrmPost(url string, data map[string]interface{}) map[string]interface{} {
	reqBody, _ := json.Marshal(data)

	req, _ := http.NewRequest("POST", os.Getenv("CRM_HOST")+url, bytes.NewBuffer([]byte(json.RawMessage(reqBody))))

	req.Header.Add("Authorization", os.Getenv("CRM_API_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()

	var respBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)
	return respBody
}
