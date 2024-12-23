package user

import (
	"go-gin/internal/httpc"
	"go-gin/util/jsonx"
)

var (
	ApiResponseSuccessCode = 200
)

// 解析返回格式固定的结构，返回结构包含code message data字段
type APIResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
	Data    any     `json:"data"`
}

var _ httpc.IResponse = (*APIResponse)(nil)

// 解析响应结构
func (r *APIResponse) Parse(b []byte) error {
	err := jsonx.Unmarshal(b, &r)
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
	dataStr, err := jsonx.Marshal(r.Data)
	if err != nil {
		return err
	}

	// 尝试将 data 字段解析为给定的结构体类型
	err = jsonx.Unmarshal(dataStr, r.Data)
	if err != nil {
		return err
	}
	return nil
}
