package validators

import (
	"github.com/go-playground/validator/v10"
)

var sample validator.Func = func(fl validator.FieldLevel) bool {
	return false
}
