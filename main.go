package main

import (
	_ "sago/internal/app"
	"sago/internal/logger"
)

func main() {
	defer logger.Flush()

	logger.Sucess("hello,sago")
}
