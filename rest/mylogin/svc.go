package mylogin

import "context"

type ILoginSvc interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
}

type LoginReq struct {
	Username string
	Pwd      string
}

type Mqtt struct {
	Customer string `json:"mqtt_customer"`
	Port     string `json:"websocket_port"`
}

type LoginResp struct {
	Lsname          string         `json:"lsname"`
	Rname           string         `json:"rname"`
	MQtt            Mqtt           `json:"mqtt_config"`
	appPointFuncArr []string       `json:"appPointFuncArr"`
	Param           map[string]any `json:"param"`
}
