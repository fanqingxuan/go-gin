package userc

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

var (
	ApiResponseSuccessCode = 200
)

type APIResponse struct {
	Code    *int        `json:"code"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

type userStruct struct {
	Payload interface{}
}

func WrapUserStruct(p interface{}) *userStruct {
	return &userStruct{
		Payload: p,
	}
}

func (uc *userStruct) ParseResponse(res interface{}, resp *resty.Response) error {
	var r APIResponse
	// str := `{"code":"200","data":{"username":"dfdfdfddf","BBBB":0},"message":"成功","requestId":"89e0813e-8be1-42b7-aedd-8022b2cdd82e"}`
	// err := json.Unmarshal([]byte(str), &r)
	err := json.Unmarshal([]byte(resp.String()), &r)
	if err != nil {
		return fmt.Errorf("error unmarshalling JSON: %w", err)
	}
	if r.Code == nil || r.Message == nil {
		return fmt.Errorf("api response format error")
	}
	if *r.Code != ApiResponseSuccessCode {
		return fmt.Errorf("api response error: %s", *r.Message)
	}

	// 将 data 字段转换为 JSON 字符串
	dataStr, err := json.Marshal(r.Data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %w", err)
	}
	// 尝试将 data 字段解析为给定的结构体类型
	err = json.Unmarshal(dataStr, uc.Payload)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %w", err)
	}
	return nil
}
