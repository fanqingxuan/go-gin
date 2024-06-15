package validators

// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"golang.org/x/text/language"
)

type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
	trans    ut.Translator
}

var _ binding.StructValidator = (*defaultValidator)(nil)

// ValidateStruct receives any kind of type, but only performed struct or pointer to struct type.
func (v *defaultValidator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		return v.ValidateStruct(value.Elem().Interface())
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(binding.SliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

// validateStruct receives struct type
func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	err := v.validate.Struct(obj)
	if err == nil {
		return err
	}
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err
	}
	// returns a map with key = namespace & value = translated error
	// NOTICE: 2 errors are returned and you'll see something surprising
	// translations are i18n aware!!!!
	// eg. '10 characters' vs '1 character'
	return v.formatValidationErrors(errs)
}

// Engine returns the underlying validator engine which powers the default
// Validator instance. This is useful if you want to register custom validations
// or struct level validations. See validator GoDoc for more info -
// https://pkg.go.dev/github.com/go-playground/validator/v10
func (v *defaultValidator) Engine() any {
	v.lazyinit()
	return v.validate
}

func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			// skip if tag key says it should be ignored
			if name == "-" {
				return ""
			}
			return name
		})
		v.trans = getZhTranslator()
		zhTrans.RegisterDefaultTranslations(v.validate, v.trans)
		v.validate.SetTagName("binding")
	})
}

// formatValidationErrors 格式化验证器返回的错误消息
func (v *defaultValidator) formatValidationErrors(errs validator.ValidationErrors) error {
	var errorMessage string
	for _, e := range errs {
		errorMessage += e.Translate(v.trans) + "\n"
	}

	return fmt.Errorf(errorMessage)
}

// getZhTranslator 获取中文翻译器
func getZhTranslator() ut.Translator {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator(language.Chinese.String())
	return trans
}
