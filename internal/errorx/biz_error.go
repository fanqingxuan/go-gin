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

func NewWithError(err error) BizError {
	return BizError{Code: ErrCodeDefaultCommon, Msg: err.Error()}
}

func (e BizError) Error() string {
	return e.Msg
}
