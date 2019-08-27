package utils

import "fmt"

//打印值
func PrintVal(s interface{}) {
	fmt.Printf("%v\n", s)
}

//打印值和类型
func PrintType(s interface{}) {
	fmt.Printf("%+v\n", s)
}
