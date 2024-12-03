package errorx

type RedisError struct {
	Msg string
}

func NewRedisError(err error) RedisError {
	return RedisError{Msg: err.Error()}
}

func (e RedisError) Error() string {
	return e.Msg
}
