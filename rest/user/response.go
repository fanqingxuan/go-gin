package user

import (
	"encoding/json"
	"go-gin/internal/httpc"
)

var (
	ApiResponseSuccessCode = 200
)

type APIResponse struct {
	Code    *int        `json:"code"`
	Message *string     `json:"message"`
	Data    interface{} `json:"data"`
}

var _ httpc.IResponse = (*APIResponse)(nil)

// 解析响应结构
func (r *APIResponse) Parse(b []byte) error {
	err := json.Unmarshal(b, &r)
	if err != nil {
		return err
	}
	return nil
}

// 验证返回格式
func (r *APIResponse) Valid() bool {
	if r.Code == nil || r.Message == nil {
		return false
	}
	return true
}

// 验证返回状态码
func (r *APIResponse) IsSuccess() bool {
	return *r.Code == ApiResponseSuccessCode
}

// 消息体
func (r *APIResponse) Msg() string {
	return *r.Message
}

// 解析数据体
func (r *APIResponse) ParseData() error {

	// 将 data 字段转换为 JSON 字符串
	dataStr, err := json.Marshal(r.Data)
	if err != nil {
		return err
	}

	// 尝试将 data 字段解析为给定的结构体类型
	err = json.Unmarshal(dataStr, r.Data)
	if err != nil {
		return err
	}
	return nil
}
