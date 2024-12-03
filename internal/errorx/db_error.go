package errorx

type DBError struct {
	Msg string
}

func (e DBError) Error() string {
	return e.Msg
}

func NewDBError(err error) DBError {
	return DBError{Msg: err.Error()}
}
