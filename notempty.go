package govalidator

import "github.com/go-playground/validator/v10"

type NotEmptyValidation struct {
}

func (n NotEmptyValidation) DefaultTpl() (tpl string, override bool) {
	return "'{0}' value required", true
}

func (n NotEmptyValidation) Tpl() (section string, key string) {
	return "govaild.tpl", "notempty"
}

func (n NotEmptyValidation) I18n(fe validator.FieldError, locale string) []string {
	return []string{fe.Field()}
}

func (n NotEmptyValidation) Tag() string {
	return "notempty"
}

func (n NotEmptyValidation) Valid(fl validator.FieldLevel) bool {
	return fl.Field().Len() > 0
}
