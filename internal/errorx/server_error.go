package errorx

import (
	"net/http"
)

type ServerError struct {
	Code int
	Msg  string
}

func NewServerError(code int) ServerError {

	msg, ok := StatusText(code)
	if !ok {
		msg = "未知错误"
	}
	return ServerError{Code: code, Msg: msg}
}

func (c ServerError) Error() string {
	return c.Msg
}

// 定义一个映射表，将 HTTP 状态码的英文描述映射成中文描述
var statusTextCN = map[int]string{
	http.StatusContinue:           "继续",
	http.StatusSwitchingProtocols: "切换协议",
	http.StatusProcessing:         "处理中",
	http.StatusEarlyHints:         "早期提示",

	http.StatusOK:                   "成功",
	http.StatusCreated:              "已创建",
	http.StatusAccepted:             "已接受",
	http.StatusNonAuthoritativeInfo: "非权威信息",
	http.StatusNoContent:            "无内容",
	http.StatusResetContent:         "重置内容",
	http.StatusPartialContent:       "部分内容",
	http.StatusMultiStatus:          "多状态",
	http.StatusAlreadyReported:      "已报告",
	http.StatusIMUsed:               "IM 已用",

	http.StatusMultipleChoices:   "多种选择",
	http.StatusMovedPermanently:  "永久移动",
	http.StatusFound:             "已找到",
	http.StatusSeeOther:          "另请参见",
	http.StatusNotModified:       "未修改",
	http.StatusUseProxy:          "使用代理",
	http.StatusTemporaryRedirect: "临时重定向",
	http.StatusPermanentRedirect: "永久重定向",

	http.StatusBadRequest:                   "错误的请求",
	http.StatusUnauthorized:                 "未经授权",
	http.StatusPaymentRequired:              "需要付款",
	http.StatusForbidden:                    "禁止访问",
	http.StatusNotFound:                     "页面不存在",
	http.StatusMethodNotAllowed:             "方法不允许",
	http.StatusNotAcceptable:                "不可接受",
	http.StatusProxyAuthRequired:            "需要代理授权",
	http.StatusRequestTimeout:               "请求超时",
	http.StatusConflict:                     "冲突",
	http.StatusGone:                         "已经不存在",
	http.StatusLengthRequired:               "需要 Content-Length",
	http.StatusPreconditionFailed:           "前置条件失败",
	http.StatusRequestEntityTooLarge:        "请求实体过大",
	http.StatusRequestURITooLong:            "请求 URI 过长",
	http.StatusUnsupportedMediaType:         "不支持的媒体类型",
	http.StatusRequestedRangeNotSatisfiable: "请求的范围无法满足",
	http.StatusExpectationFailed:            "预期失败",
	http.StatusTeapot:                       "茶壶",
	http.StatusMisdirectedRequest:           "错误的请求方向",
	http.StatusUnprocessableEntity:          "无法处理的实体",
	http.StatusLocked:                       "已锁定",
	http.StatusFailedDependency:             "依赖失败",
	http.StatusTooEarly:                     "请求过早",
	http.StatusUpgradeRequired:              "需要升级",
	http.StatusPreconditionRequired:         "需要预加载",
	http.StatusTooManyRequests:              "请求过多",
	http.StatusRequestHeaderFieldsTooLarge:  "请求头过大",
	http.StatusUnavailableForLegalReasons:   "因法律原因不可用",

	http.StatusInternalServerError:           "服务器内部错误",
	http.StatusNotImplemented:                "未实现",
	http.StatusBadGateway:                    "错误的网关",
	http.StatusServiceUnavailable:            "服务不可用",
	http.StatusGatewayTimeout:                "网关超时",
	http.StatusHTTPVersionNotSupported:       "HTTP 版本不支持",
	http.StatusVariantAlsoNegotiates:         "变种也谈判",
	http.StatusInsufficientStorage:           "存储空间不足",
	http.StatusLoopDetected:                  "检测到循环",
	http.StatusNotExtended:                   "未扩展",
	http.StatusNetworkAuthenticationRequired: "需要网络认证",
}

func StatusText(code int) (string, bool) {
	text, ok := statusTextCN[code]
	if ok {
		return text, true
	}
	return "", false
}
