package errorx

type RedisError struct {
	Msg string
}

func NewRedisError(err error) error {
	if err == nil {
		return nil
	}
	return RedisError{Msg: err.Error()}
}

func (e RedisError) Error() string {
	return e.Msg
}
