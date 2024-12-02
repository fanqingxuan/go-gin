package errorx

import "net/http"

const (
	ErrCodeDefaultCommon  = 10000 // 默认通用错误码
	ErrCodeValidateFailed = 10001 // 验证失败
)

var (
	ErrThirdAPIConnectFailed          = New(11001, "第三方接口连接失败")
	ErrThirdAPIContentNoContentFailed = New(11002, "第三方接口返回内容为空")
	ErrThirdAPIContentParseFailed     = New(11003, "第三方接口响应内容解析失败")
	ErrThirdAPICallFormatFailed       = New(11004, "第三方接口返回格式不正确")
	ErrThirdAPIDataParseFailed        = New(11005, "第三方接口解析data失败")
	ErrThirdAPIBusinessFailed         = New(11006, "第三方接口业务上的错误")
)

var ErrRedisOperateFailed = NewRedisError(12001, "服务器内部错误")
var ErrDBOperateFailed = NewDBError(12002, "服务器内部错误")

// http错误
var (
	ErrMethodNotAllowed    = NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = NewServerError(http.StatusNotFound)
	ErrInternalServerError = NewServerError(http.StatusInternalServerError)
)
