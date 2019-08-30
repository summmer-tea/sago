package sago

import (
	"bufio"
	"encoding/json"
	"fmt"
	utils "gitee.com/xiawucha365/sago/internal/tool"
	"github.com/parnurzeal/gorequest"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
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

//文件相关

func (t *Tooler) CreateDir(dir string) {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

//WriteFile
func (t *Tooler) WriteFile(filename string, data string) bool {

	if t.Exist(filename) {
		if err := ioutil.WriteFile(filename, []byte(data), 0666); err != nil {
			panic(err)
		}
	} else {
		if file, err := os.Create(filename); err != nil {
			panic(err)
		} else {
			if _, err := file.Write([]byte(data)); err != nil {
				panic(err)
			}
			//关闭句柄
			defer file.Close()
		}

	}

	return true
}

func (t *Tooler) ReadFile(fileName string) (data []byte) {
	if !t.Exist(fileName) {
		if file, err := os.Create(fileName); err != nil {
			panic(err)
		} else {
			file.Close()
			return []byte("")
		}
	}
	//读取文件
	if file, err := os.Open(fileName); err != nil {
		panic(err)
	} else {
		defer file.Close()
		content, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		} else {
			return content
		}
	}

}

func (t *Tooler) Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func GetCurrentPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (t *Tooler) GetExecDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

//获取项目根目录
func (t *Tooler) GetRootPath() string {
	currPath := GetCurrentPath()
	index := strings.LastIndex(currPath, "hotel_scripts")
	return string([]rune(currPath)[:index+len("hotel_scripts")])
}

// SelfPath gets compiled executable file absolute path
func SelfPath() string {
	path, _ := filepath.Abs(os.Args[0])
	return path
}

// SelfDir gets compiled executable file directory
func (t *Tooler) SelfDir() string {
	return filepath.Dir(SelfPath())
}

// FileExists reports whether the named file or directory exists.
// 判断所给路径文件/文件夹是否存在
func (t *Tooler) FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// GrepFile like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func (t *Tooler) GrepFile(patten string, filename string) (lines []string, err error) {
	re, err := regexp.Compile(patten)
	if err != nil {
		return
	}

	fd, err := os.Open(filename)
	if err != nil {
		return
	}
	lines = make([]string, 0)
	reader := bufio.NewReader(fd)
	prefix := ""
	var isLongLine bool
	for {
		byteLine, isPrefix, er := reader.ReadLine()
		if er != nil && er != io.EOF {
			return nil, er
		}
		if er == io.EOF {
			break
		}
		line := string(byteLine)
		if isPrefix {
			prefix += line
			continue
		} else {
			isLongLine = true
		}

		line = prefix + line
		if isLongLine {
			prefix = ""
		}
		if re.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines, nil
}
