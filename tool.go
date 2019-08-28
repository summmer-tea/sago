package sago

import utils "gitee.com/xiawucha365/sago/internal/tool"

var Tool *Tooler

type Tooler struct{}

//文件相关
func (t *Tooler) JsonEncode(v ...interface{}) (string, error) {
	return utils.JsonEncode(v)
}

func (t *Tooler) JsonDecode(data string, v interface{}) error {
	return utils.JsonDecode(data, v)
}

//http
func (t *Tooler) RequestGet(url string) string {
	return utils.Get(url)
}

func (t *Tooler) RequestPostJson(url string) string {
	return utils.Get(url)
}

func (t *Tooler) RequestPostForm(url string) string {
	return utils.Get(url)
}
