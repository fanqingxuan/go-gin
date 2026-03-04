package mylogin

import (
	"go-gin/internal/httpc"
	"go-gin/util/jsonx"
)

var (
	ApiResponseSuccessCode = true
)

// 解析返回格式不固定，但是success msg两个标准字段，其它业务字段格式不固定
type APIResponse struct {
	Code    *bool   `json:"success"`
	Message *string `json:"msg"`
	Data    any
}

var _ httpc.IResponseNonStandard = (*APIResponse)(nil)

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
func (r *APIResponse) ParseData(b []byte) error {
	return jsonx.Unmarshal(b, r.Data)
}
