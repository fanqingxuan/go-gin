package validators

import (
	"go-gin/internal/errorx"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Trans ut.Translator

func RegisterValidators() {
	binding.Validator = &defaultValidator{}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sample", sample)
	}
}

func Validate(obj any) error {
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		return errorx.NewDefault(err.Error())
	}
	return nil
}
