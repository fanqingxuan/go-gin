package errorx

type DBError struct {
	Code int
	Msg  string
}

func (e DBError) Error() string {
	return e.Msg
}

func NewDBError(code int, msg string) DBError {
	return DBError{Code: code, Msg: msg}
}
