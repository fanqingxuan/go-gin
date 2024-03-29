package validators

import (
	"go-gin/internal/errorx"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
)

var Trans ut.Translator

func RegisterValidators() {
	binding.Validator = &defaultValidator{}

}

func Validate(obj any) error {
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		return errorx.NewDefault(err.Error())
	}
	return nil
}
