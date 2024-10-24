package validators

import (
	"github.com/dchest/captcha"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
)

type CacheCodeValidator struct {
	ginValidator.BaseValidator
}

func (c *CacheCodeValidator) TagName() string {
	return "cacheCode"
}

// Messages 规则错误提示信息
func (c *CacheCodeValidator) Messages() string {
	return "{0} 验证码错误"
}

// Validator 规则验证逻辑
func (c *CacheCodeValidator) Validator(fl validator.FieldLevel) bool {
	keyId := fl.Param()
	return captcha.VerifyString(fl.Parent().FieldByName(keyId).String(), fl.Field().String())
}
