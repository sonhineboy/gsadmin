package initialize

import (
	"github.com/sonhineboy/gsadmin/service/app/validators"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
)

func InitValidator() *ginValidator.CustomValidatorManager {
	return initCustomValidator(initTrans())
}

func initTrans() *ginValidator.Trans {
	tran := ginValidator.NewDefaultTrans()
	err := tran.SetUp()
	if err != nil {
		panic(err)
	}
	return tran
}

func initCustomValidator(tran *ginValidator.Trans) *ginValidator.CustomValidatorManager {
	customValidator := ginValidator.NewCustomValidatorManager(make(map[string]ginValidator.CustomValidator), tran.GetValidate(), tran.GetTrans())
	customValidator.Adds(
		new(validators.CacheCodeValidator),
	)
	customValidator.RegisterToValidate()
	return customValidator
}
