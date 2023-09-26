package validatorutil

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

// MyValidate 参数校验方法，由于自带错误提示看起来对前端不是太友好，简单处理一下
func MyValidate(data interface{}) (err error) {
	err = validator.New().Struct(data)
	if err == nil {
		return
	}
	errMsg := "参数错误："
	for _, v := range err.(validator.ValidationErrors) {
		errMsg += v.StructField() + " " + v.Tag() + "，"
	}
	errMsg = strings.TrimRight(errMsg, "，")
	return errors.New(errMsg)
}
