package audit

import (
	"fmt"
	"os"
	"regexp"
)

func Audit(data string, rules string) bool {
	reg1 := regexp.MustCompile(`(?i)(` + rules + `)`)
	// 不匹配大小写
	if reg1 == nil {
		fmt.Println("regexp err")
		os.Exit(1)
	}
	// 正则出错

	result1 := reg1.FindAllStringSubmatch(data, -1)
	return result1 != nil
}

// 基本匹配规则：正则匹配成功返回true，否则返回false
