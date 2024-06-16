package errorx

type BizError struct {
	Code int
	Msg  string
}

func New(code int, msg string) BizError {
	return BizError{Code: code, Msg: msg}
}

func (c BizError) Error() string {
	return c.Msg
}
