package errorx

type RedisError struct {
	Msg string
}

func TryToRedisError(err error) error {
	if err == nil {
		return nil
	}
	return RedisError{Msg: err.Error()}
}

func (e RedisError) Error() string {
	return e.Msg
}
