package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadminValidator/ginValidator"
)

type DemoValidator struct {
	ginValidator.BaseValidator
}

func (d *DemoValidator) TagName() string {
	return "Demo"
}

func (d *DemoValidator) Messages() string {
    //This is error message
	return ""
}

func (d *DemoValidator) Validator(fl validator.FieldLevel) bool {
	//To Do .....
	return true
}
