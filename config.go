package sago

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	// 单例
	GConfig *Config
)

// 程序配置
type Config struct {
	Common ComConfig
	Mysql  MysqlConfig
	Redis  RedisConfig
}

type MysqlConfig struct {
	Addr     string
	Username string
	Password string
	Charset  string
	Dbname   string
}

type RedisConfig struct {
	Addr     string
	Password string
}

type ComConfig struct {
	Appdir    string // 应用根目录
	ConfDir   string // 配置文件目录
	Logdir    string //日志目录
	VersionId int    // 版本号
	Debug     bool
	Env       string
}

// 加载配置
func InitConfig(filename string) (err error) {
	var (
		content []byte
		conf    Config
	)
	// 1, 把配置文件读进来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return err
	}

	if _, err := toml.Decode(string(content), &conf); err != nil {
		return err
	}

	GConfig = &conf

	return
}
