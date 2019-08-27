package utils

import (
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

func Get(url string) string {
	_, body, _ := gorequest.New().Get(url).End()
	return body
}

/*
	json提交
	json_input := map[string]interface{}{
		"name": "backy",
		"species": "dog",
	}

	json_input := `{"name":"backy", "species":"dog"}`

*/
func PostJson(url string, json_input interface{}) (string, *http.Response, []error) {
	if resp, body, errs := gorequest.New().Post(url).Type("json").Send(json_input).
		Timeout(30 * time.Second).End(); errs != nil {
		return "", resp, errs
	} else {
		return body, resp, nil
	}
}

//表单提交
func PostForm(url string, json_input interface{}) (string, *http.Response, []error) {
	if resp, body, errs := gorequest.New().
		Post(url).
		//Set("Content-Type","application/x-www-form-urlencoded").
		Type("form").
		Send(json_input).
		Timeout(30 * time.Second).End(); errs != nil {
		return "", resp, errs
	} else {
		return body, resp, nil
	}

}

func printBody(resp gorequest.Response, body string, errs []error) {
	if errs != nil {
		//app.G_log.Warn("statusCode:", resp.Status, errs)
	}
}
