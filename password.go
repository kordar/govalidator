package govalidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PasswordValidation struct {
}

func (n PasswordValidation) Tag() string {
	//TODO implement me
	return "password"
}

func (n PasswordValidation) Valid(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^[a-zA-Z0-9_-]{8,16}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func (n PasswordValidation) I18n() (section string, key string) {
	//TODO implement me
	return "validation.default", "password"
}
