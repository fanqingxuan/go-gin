package svc

import (
	"go-gin/config"
	"go-gin/svc/redisx"
	"go-gin/svc/sqlx"
)

type ServiceContext struct {
	// Config config.Config
	Redis *redisx.Redisx
	DB    sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("mysql", c.Mysql)
	return &ServiceContext{
		// Config: c,
		Redis: redisx.New(),
		DB:    conn,
	}
}
