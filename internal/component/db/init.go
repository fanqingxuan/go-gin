package db

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	conf     Config
)

type Config struct {
	DSN          string `yaml:"dsn"`
	MaxOpenConns int    `yaml:"max-open-conn"`
	MaxIdleConns int    `yaml:"max-idle-conn"`
	LogLevel     string `yaml:"log-level"`
}

func InitConfig(c Config) {
	conf = c
}

func Init() {
	err := Connect()
	if err != nil {
		panic(err)
	}
}

func IsConnected() bool {
	return instance != nil
}

func Connect() (err error) {
	instance, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       conf.DSN, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: &DBLog{
			LogLevel: ParseLevel(conf.LogLevel),
		},
	})
	if err != nil {
		instance = nil
		return err
	}

	sqlDB, err := instance.DB()
	if err != nil {
		instance = nil
		return err
	}

	if err = sqlDB.Ping(); err != nil {
		instance = nil
		return err
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	return
}

// NewModel 创建一个新的 Model 实例用于链式操作（GoFrame 风格）
func NewModel(ctx context.Context, table string) *Model {
	return newModel(ctx, table)
}
