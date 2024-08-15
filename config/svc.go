package config

import (
	"go-gin/rest/login"
	"go-gin/rest/mylogin"
	"go-gin/rest/user"
)

type SvcConfig struct {
	UserSvcUrl  string `yaml:"user_url"`
	LoginSvcUrl string `yaml:"login_url"`
}

func InitSvc() {
	svcConfig := instance.Svc
	// 初始化第三方请求服务
	user.Init(svcConfig.UserSvcUrl)
	login.Init(svcConfig.LoginSvcUrl)
	mylogin.Init(svcConfig.LoginSvcUrl)
}
