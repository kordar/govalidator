package govalidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PasswordValidation struct {
	Rules map[string]string
}

func (n PasswordValidation) DefaultTpl() (tpl string, override bool) {
	return "password format error", true
}

func (n PasswordValidation) Tpl() (section string, key string) {
	//TODO implement me
	return "govaild.tpl", "password"
}

func (n PasswordValidation) I18n(fe validator.FieldError, locale string) []string {
	return nil
}

func (n PasswordValidation) Tag() string {
	//TODO implement me
	return "password"
}

func (n PasswordValidation) Valid(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^[a-zA-Z0-9_-]{8,16}$"
	if n.Rules[fl.Param()] != "" {
		regular = n.Rules[fl.Param()]
	}
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
