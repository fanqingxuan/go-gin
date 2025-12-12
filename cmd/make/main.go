// 代码生成工具集
// 用法:
//   go run cmd/make/main.go make:enum                    生成枚举代码
//   go run cmd/make/main.go make:dao                     生成所有表的 dao/entity/do
//   go run cmd/make/main.go make:dao -t user,order       生成指定表的 dao/entity/do
//   go run cmd/make/main.go make:migration <name>        生成迁移文件
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	cmd := os.Args[1]
	args := os.Args[2:]

	switch cmd {
	case "make:enum":
		runEnum()
	case "make:dao":
		runDao(args)
	case "make:migration":
		runMigration(args)
	default:
		fmt.Printf("未知命令: %s\n\n", cmd)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`代码生成工具集

用法:
  go run cmd/make/main.go <command> [options]

命令:
  make:enum                    扫描 const/enum 目录，生成枚举方法
  make:dao                     扫描数据库，生成 entity/do/dao 代码
  make:dao -t user,order       生成指定表的代码
  make:migration <name>        生成迁移文件

示例:
  go run cmd/make/main.go make:enum
  go run cmd/make/main.go make:dao -f .env
  go run cmd/make/main.go make:dao -f .env -t user,order
  go run cmd/make/main.go make:migration create_orders`)
}
