package errorx

import "net/http"

const (
	ErrCodeDefaultCommon   = 10000 // 默认通用错误码
	ErrCodeValidateFailed  = 10001 // 验证失败
	ErrCodeDBOperateFailed = 10002 // 数据库操作错误
)

const (
	ErrCodeThirdAPIConnectFailed          = 11001 // 第三方接口连接失败
	ErrCodeThirdAPIContentNoContentFailed = 11002 // 第三方接口返回内容为空
	ErrCodeThirdAPIContentParseFailed     = 11003 // 第三方接口响应内容解析失败
	ErrCodeThirdAPICallFormatFailed       = 11004 // 第三方接口返回格式不正确
	ErrCodeThirdAPIDataParseFailed        = 11005 // 第三方接口解析data失败
	ErrCodeThirdAPIBusinessFailed         = 11006 // 第三方接口业务上的错误

)

// http错误
var (
	ErrMethodNotAllowed    = NewServerError(http.StatusMethodNotAllowed)
	ErrNoRoute             = NewServerError(http.StatusNotFound)
	ErrInternalServerError = NewServerError(http.StatusInternalServerError)
)
