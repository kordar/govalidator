package govalidator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type PhoneValidation struct {
	Rules map[string]string
}

func (p PhoneValidation) DefaultTpl() (tpl string, override bool) {
	return "wrong phone number", true
}

func (p PhoneValidation) Tpl() (section string, key string) {
	return "govaild.tpl", "phone"
}

func (p PhoneValidation) I18n(fe validator.FieldError, locale string) []string {
	return nil
}

func (p PhoneValidation) Tag() string {
	return "phone"
}

func (p PhoneValidation) Valid(fl validator.FieldLevel) bool {
	mobileNum := fl.Field().String()
	fl.StructFieldName()
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	if p.Rules[fl.Param()] != "" {
		regular = p.Rules[fl.Param()]
	}
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
