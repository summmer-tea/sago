package sago

import (
	"gitee.com/xiawucha365/sago/internal/logger"
	"github.com/go-xorm/xorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"time"
)

var (
	once sync.Once
	//通用模式,可以调用xorm底层方法
	MysqlEngine *DbEngine
)

type DbEngine = xorm.Engine

func CreateMysqlDialect() *DbEngine {
	once.Do(func() {
		var err error
		if MysqlEngine, err = xorm.NewEngine("mysql",
			GConfig.Mysql.Username+":"+GConfig.Mysql.Password+
				"@tcp("+GConfig.Mysql.Addr+")/"+GConfig.Mysql.Dbname+"?charset="+
				GConfig.Mysql.Charset+"&parseTime=True&loc=Local"); err != nil {
			logger.Error("mysqlconn", err)
		}

		//mysqlConn.ShowSQL(true)
		//engine, err = xorm.NewEngine("mysql", comm.G_config.Mysql.Username+":"+comm.G_config.Mysql.Password+"@/"+comm.G_config.Mysql.Dbname+"?charset=utf8")
		//连接数默认设置
		//连接数默认设置

		MysqlEngine.SetMaxOpenConns(20)
		MysqlEngine.SetMaxIdleConns(10)
		MysqlEngine.SetConnMaxLifetime(time.Second * 20)
	})
	return MysqlEngine
}
