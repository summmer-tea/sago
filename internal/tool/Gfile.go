package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func CreateDir(dir string) {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		fmt.Printf("%s", err)
	}
}

//WriteFile
func WriteFile(filename string, data string) bool {

	if Exist(filename) {
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

func ReadFile(fileName string) (data []byte) {
	if !Exist(fileName) {
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

func Exist(filename string) bool {
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

func GetExecDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

//获取项目根目录
func GetRootPath() string {
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
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// FileExists reports whether the named file or directory exists.
// 判断所给路径文件/文件夹是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// SearchFile Search a file in paths.
// this is often used in search config file in /etc ~/
func SearchFile(filename string, paths ...string) (fullpath string, err error) {
	for _, path := range paths {
		if fullpath = filepath.Join(path, filename); FileExists(fullpath) {
			return
		}
	}
	err = errors.New(fullpath + " not found in paths")
	return
}

// GrepFile like command grep -E
// for example: GrepFile(`^hello`, "hello.txt")
// \n is striped while read
func GrepFile(patten string, filename string) (lines []string, err error) {
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
