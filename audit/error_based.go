package audit

import (
	"fmt"
	"main/pre"
	"regexp"
)

func Error_based_detection(data string) {
	// 用于探测报错注入
	poc := []string{"updatexml", "exp", "extractvalue", "floor"}

	res := pre.Concat(poc, "|")
	reg1 := regexp.MustCompile(`(?i)(` + res + `)`)
	// 不匹配大小写
	if reg1 == nil {
		fmt.Println("regexp err")
		return
	}
	// 正则出错

	result1 := reg1.FindAllStringSubmatch(data, -1)
	if result1 != nil {
		fmt.Println("[+] Potential Error-Based Injection: ", data)
		// 匹配到关键字
	}
}
