package svc

import "go-gin/svc/redisx"

type ServiceContext struct {
	// Config config.Config
	Redis *redisx.Redisx
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		// Config: c,
		Redis: redisx.New(),
	}
}
