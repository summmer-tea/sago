package sago

import (
	"flag"
	"gitee.com/xiawucha365/sago/internal/logger"
	tool "gitee.com/xiawucha365/sago/internal/tool"
	"github.com/garyburd/redigo/redis"
)

const (
	GENV_CONST_DEV  = "dev"
	GENV_CONST_PROD = "prod"
)

var (
	GMysql *DbEngine
	GModel *DbDialect
	GRedis redis.Conn
	Env    string
	Debug  bool
)

func init() {
	initEnv()
	initMysql()
	initRedis()

	defer func() {
		//日志缓冲完成
		logger.Flush()
	}()
}

// InitEnv 初始化
func initEnv() {

	var etc_file string
	etc_dir := tool.GetCurrentPath()

	flag.StringVar(&Env, "env", "dev", "请选择环境:test<测试环境> prod<线上环境>")

	flag.BoolVar(&Debug, "debug", false, "是否输出详细调试信息")

	flag.Parse()

	//启动命令
	if Env == GENV_CONST_PROD {
		etc_file = etc_dir + "/etc/prod.toml"
	}
	if Env == GENV_CONST_DEV {
		etc_file = "/Users/mfw/Documents/data/go/src/hotel_scripts/console/spider_worker/etc/dev.toml"
	}

	if err := InitConfig(etc_file); err != nil {
		logger.Error(err)
	}

	GConfig.Common.Env = Env
	GConfig.Common.Debug = Debug

	initLog()

	return
}

func initMysql() {
	//通用模式
	GMysql = CreateMysqlDialect()
	//自定义封装模式
	GModel = MysqlDialect
}

func initRedis() {
	if GConfig.Redis.Addr == "" {
		return
	}
	GRedis = createRedisDialect()
}

//日志初始化
func initLog() {
	logger.Init(GConfig.Common.Logdir)
}
