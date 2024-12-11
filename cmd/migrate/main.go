package main

import (
	"flag"
	filex "go-gin/internal/file"
	"go-gin/internal/migration"

	"github.com/labstack/gommon/color"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// 导入所有迁移文件
	_ "go-gin/migrations/ddl"
	_ "go-gin/migrations/dml"
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
	// 数据库连接配置
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      c.DB.DSN, // DSN data source name
		DisableDatetimePrecision: true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:   true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:  true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列

	}), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		color.Printf(color.Red("Database connection failed: %v\n"), err)
		return
	}
	// 设置数据库连接
	migration.SetDB(db)

	// 执行迁移
	if err := migration.GetManager().Run(); err != nil {
		return
	}
}
