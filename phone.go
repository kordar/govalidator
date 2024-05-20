package govalidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PhoneValidation struct {
}

func (p PhoneValidation) Tag() string {
	return "phone"
}

func (p PhoneValidation) Valid(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func (p PhoneValidation) I18n() (section string, key string) {
	//TODO implement me
	return "validation.default", "phone"
}
