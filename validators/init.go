package validators

import (
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
