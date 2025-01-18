package errorx

import "net/http"

const (
	ErrCodeDefault            = 10000 // 默认通用错误码
	ErrCodeBizDefault         = 10001 // 业务默认错误
	ErrCodeValidateFailed     = 10002 // 验证失败
	ErrCodeDBOperateFailed    = 10003 // 数据库操作失败
	ErrCodeRedisOperateFailed = 10004 // redis操作失败
)

var (
	ErrThirdAPIConnectFailed          = New(11001, "第三方接口连接失败")
	ErrThirdAPIContentNoContentFailed = New(11002, "第三方接口返回内容为空")
	ErrThirdAPIContentParseFailed     = New(11003, "第三方接口响应内容解析失败")
	ErrThirdAPICallFormatFailed       = New(11004, "第三方接口返回格式不正确")
	ErrThirdAPIDataParseFailed        = New(11005, "第三方接口解析data失败")
	ErrThirdAPIBusinessFailed         = New(11006, "第三方接口业务上的错误")
)

// http错误
var (
	ErrorBadRequest        = NewServerError(http.StatusBadRequest)
	ErrorUnauthorized      = NewServerError(http.StatusUnauthorized)
	ErrorForbidden         = NewServerError(http.StatusForbidden)
	ErrMethodNotAllowed    = NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = NewServerError(http.StatusNotFound)
	ErrInternalServerError = NewServerError(http.StatusInternalServerError)
)
