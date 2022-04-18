package main

import (
	"main/audit"
	"main/pre"
)

func main() {
	file := `test.txt`
	data := pre.Readlines(file)
	//fmt.Println(data)
	for i := 0; i < len(data); i++ {
		audit.Error_based_detection(data[i])
		audit.Time_based_bool_detection(data[i])
	}
	// 按切片解析输入
}
