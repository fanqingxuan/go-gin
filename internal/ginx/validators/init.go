package validators

import (
	"github.com/gin-gonic/gin/binding"
)

func Init() {
	binding.Validator = &defaultValidator{}
}

func Validate(obj any) error {
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		return err
	}
	return nil
}
