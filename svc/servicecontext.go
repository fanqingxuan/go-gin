package svc

import (
	"go-gin/config"
	"go-gin/svc/redisx"
	"go-gin/svc/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type ServiceContext struct {
	// Config config.Config
	Redis *redisx.Redisx
	DB    sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		// Config: c,
		Redis: redisx.New(),
		DB:    sqlx.NewSqlConn("mysql", c.Mysql.DataSource),
	}
}
