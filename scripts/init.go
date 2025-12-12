// Package scripts 一次性脚本
// 用法: go run scripts/init.go scripts/xxx.go -f .env
package scripts

import (
	"flag"

	"go-gin/config"
	"go-gin/internal/component/db"
)

var ConfigFile = flag.String("f", "./.env", "the config file")

// Init 初始化配置和数据库连接
func Init() {
	flag.Parse()
	config.Init(*ConfigFile)

	dbConf := config.GetDbConf()
	db.InitConfig(dbConf)
	db.Init()
}
