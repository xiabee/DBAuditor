package main

import "main/audit"

func main() {
	sql_file := `test.txt`
	// 待审计文件
	rules_dir := `./rules/sql_injection`
	// 审计规则
	audit.Sql_injection_detection(sql_file, rules_dir)
}
