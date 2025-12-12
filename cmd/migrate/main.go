// 数据库迁移工具
// 用法: go run cmd/migrate/main.go -f .env
package main

import (
	"flag"

	filex "go-gin/internal/file"
	"go-gin/internal/migration"

	"github.com/labstack/gommon/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// 导入迁移文件
	_ "go-gin/migration"
)

var configFile = flag.String("f", "./.env", "the config file")

type DBConfig struct {
	DSN string `yaml:"dsn"`
}

type Config struct {
	DB DBConfig `yaml:"db"`
}

func main() {
	flag.Parse()
	color.Enable()

	var c Config
	filex.MustLoad(*configFile, &c)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      c.DB.DSN,
		DisableDatetimePrecision: true,
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,
	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		color.Printf(color.Red("Database connection failed: %v\n"), err)
		return
	}

	migration.SetDB(db)

	if err := migration.GetManager().Run(); err != nil {
		return
	}
}
