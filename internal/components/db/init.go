package db

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var conn *gorm.DB

type Config struct {
	DSN          string
	MaxOpenConns int
	MaxIdleConns int
}

func Init(c Config) {
	var err error

	conn, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       c.DSN, // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: &my_log{},
	})

	sqlDB, err := conn.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

}

func WithContext(ctx context.Context) *gorm.DB {
	return conn.WithContext(ctx)
}
