# DBAuditor

`SQL`数据库审计系统，目前支持`SQL`注入攻击审计



### 环境配置

```bash
sudo apt install golang
```

### 

### 运行方式

将待审计语句填入`test.txt`中，然后运行主程序：

* 直接运行：

```bash
go run main.go
```

* 编译运行：

```bash
go build main.go
./main
```



### 主要目录结构

```bash
.
├── audit
│   ├── audit.go
│   └── sql_injection.go
├── go.mod
├── main.go
├── pre
│   └── pre.go
├── rules
│   ├── sensitive_db
│   └── sql_injection
└── test.txt
```

* `aduit`：审计功能代码目录

* `pre`：辅助功能代码目录

* `rules`：审计规则目录

* `main.go`：主程序入口

* `test.txt`：待审计文件
