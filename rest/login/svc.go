package login

import "context"

type ILoginSvc interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
}

type LoginReq struct {
	Username string
	Pwd      string
}

type LoginResp struct {
	AppGoodInDepotType          string `json:"AppGoodInDepotType"`
	App_estimate_his_week_count string `json:"App_estimate_his_week_count"`
	AppPredictdishOrder         string `json:"AppPredictdishOrder"`
}
