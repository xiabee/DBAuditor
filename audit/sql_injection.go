package audit

import (
	"fmt"
	"main/pre"
)

func Sql_injection_detection(sql_file string, rules_dir string) bool {

	if !pre.HasSuffix(rules_dir, "/") {
		rules_dir += "/"
	}
	// 目录拼接预处理

	sql_data := pre.Readlines(sql_file)
	rules_file := pre.Get_files(rules_dir)

	for _, current_file := range rules_file {
		// 通过每个规则文件进行审计
		rules_data := pre.Readlines(rules_dir + current_file)
		// 拼接相对目录：main函数执行目录

		for _, current_sql := range sql_data {
			// 每一条数据
			for _, current_rule := range rules_data {
				// 每一条审计规则
				if Audit(current_sql, current_rule) {
					fmt.Println("[+]" + current_file + " :" + current_sql)
					break
					// 判定为注入就不再继续审计，直接审计下一条sql语句
				}

			}
		}

	}

	return true
}
