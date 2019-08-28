package sago

import (
	"gitee.com/xiawucha365/sago/internal/logger"
)

var (
	// 单例
	Log *Logger
)

// 程序配置
type Logger struct{}

func (l *Logger) Warn(v ...interface{}) {
	logger.Warn(v)
}

func (l *Logger) Error(v ...interface{}) {
	logger.Error(v)
}

func (l *Logger) Info(v ...interface{}) {
	logger.Info(v)
}

func (l *Logger) Sucess(v ...interface{}) {
	logger.Sucess(v)
}

func (l *Logger) Flush(v ...interface{}) {
	logger.Flush()
}
