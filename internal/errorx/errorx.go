package errorx

const DefaultErrorCode = 1001

type MYError struct {
	Code int
	Msg  string
}

func NewDefault(msg string) MYError {
	return New(DefaultErrorCode, msg)
}

func New(code int, msg string) MYError {
	return MYError{Code: code, Msg: msg}
}

func NewWithError(err error) MYError {
	return NewDefault(err.Error())
}

func (c MYError) Error() string {
	return c.Msg
}
