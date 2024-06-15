package validators

import (
	"go-gin/internal/errorx"

	"github.com/gin-gonic/gin/binding"
)

func Init() {
	binding.Validator = &defaultValidator{}

}

func Validate(obj any) error {
	if err := binding.Validator.ValidateStruct(obj); err != nil {
		return errorx.NewDefault(err.Error())
	}
	return nil
}
