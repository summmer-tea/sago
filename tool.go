package sago

import (
	"encoding/json"
	"fmt"
	utils "gitee.com/xiawucha365/sago/internal/tool"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strconv"
	"time"
)

var Tool *Tooler

type Tooler struct{}

//文件相关
func (t *Tooler) JsonEncode(v interface{}) (string, error) {

	if by, error := json.Marshal(v); error != nil {
		return "", error
	} else {
		return string(by), nil
	}
}

func (t *Tooler) JsonDecode(data string, v interface{}) error {
	if error := json.Unmarshal([]byte(data), v); error != nil {
		return error
	}
	return nil
}

//http
func (t *Tooler) RequestGet(url string) string {
	return utils.Get(url)
}

func (t *Tooler) RequestPostJson(url string, json_input interface{}) (string, *http.Response, []error) {
	if resp, body, errs := gorequest.New().Post(url).Type("json").Send(json_input).
		Timeout(30 * time.Second).End(); errs != nil {
		return "", resp, errs
	} else {
		return body, resp, nil
	}
}

func (t *Tooler) RequestPostForm(url string, json_input interface{}) (string, *http.Response, []error) {
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

func (t *Tooler) ConvStr2Int(str string) int {
	if intb, err := strconv.Atoi(str); err != nil {
		fmt.Println("ConvStr2Int:error")
		panic(err)
	} else {
		return intb
	}
}

func (t *Tooler) ConvInt2Str(str int) string {
	return strconv.Itoa(str)
}

//浮点保留2位小数
func (t *Tooler) MathDecimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

//打印值
func (t *Tooler) PrintVal(s interface{}) {
	fmt.Printf("%v\n", s)
}

//打印值和类型
func (t *Tooler) PrintType(s interface{}) {
	fmt.Printf("%+v\n", s)
}
