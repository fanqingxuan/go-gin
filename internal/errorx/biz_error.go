package errorx

type BizError struct {
	Code int
	Msg  string
}

func New(code int, msg string) BizError {
	return BizError{Code: code, Msg: msg}
}

func NewDefault(msg string) BizError {
	return BizError{Code: ErrCodeDefaultCommon, Msg: msg}
}

func NewWithError(code int, err error) BizError {
	return BizError{Code: code, Msg: err.Error()}
}

func (c BizError) Error() string {
	return c.Msg
}
