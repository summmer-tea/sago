package utils

import (
	"fmt"
	"strconv"
)

//浮点保留2位小数
func MathDecimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}
