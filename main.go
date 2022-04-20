package main

import "main/audit"

func main() {
	sql_file := `test.txt`
	rules_dir := `./rules/sql_injection`
	audit.Sql_injection_detection(sql_file, rules_dir)
}
