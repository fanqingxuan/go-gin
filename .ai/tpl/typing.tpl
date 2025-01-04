package typing

type XxxReq struct {
	Name   string    `form:"name" binding:"required" label:"姓名"`
}

type XxxResp struct {
	Message string `json:"message"`
}