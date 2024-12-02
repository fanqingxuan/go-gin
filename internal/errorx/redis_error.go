package errorx

type RedisError struct {
	Code int
	Msg  string
}

func NewRedisError(code int, msg string) RedisError {
	return RedisError{Code: code, Msg: msg}
}

func (e RedisError) Error() string {
	return e.Msg
}
