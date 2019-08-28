package sago

import (
	"gitee.com/xiawucha365/sago/internal/logger"
)

func Warn(v ...interface{}) {
	logger.Warn(v)
}

func Error(v ...interface{}) {
	logger.Error(v)
}

func Info(v ...interface{}) {
	logger.Info(v)
}

func Sucess(v ...interface{}) {
	logger.Sucess(v)
}
