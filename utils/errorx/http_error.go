package errorx

type HttpError struct {
	Code int
	Msg  string
}

func NewHHttpError(code int, msg string) HttpError {
	if code < 100 || code > 511 {
		// 服务端错误不合法
		panic("server error code not illegal")
	}
	return HttpError{Code: code, Msg: msg}
}

func (c HttpError) Error() string {
	return c.Msg
}
