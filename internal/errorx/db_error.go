package errorx

import (
	"errors"

	"gorm.io/gorm"
)

type DBError struct {
	Msg string
}

func (e DBError) Error() string {
	return e.Msg
}

func NewDBError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return DBError{Msg: err.Error()}
}
