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
	return jsonx.Unmarshal(b, &r)
}

// 验证返回格式
func (r *APIResponse) Valid() bool {
	return r.Code != nil && r.Message != nil
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
	dataStr, err := jsonx.Marshal(r.Data)
	if err != nil {
		return err
	}
	return jsonx.Unmarshal(dataStr, r.Data)
}
