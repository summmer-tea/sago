package utils

import (
	"fmt"
	"strconv"
)

func ConvStr2Int(str string) int {
	if intb, err := strconv.Atoi(str); err != nil {
		fmt.Println("ConvStr2Int:error")
		panic(err)
	} else {
		return intb
	}
}

func ConvInt2Str(str int) string {
	return strconv.Itoa(str)
}
