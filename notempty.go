package govalidator

import "github.com/go-playground/validator/v10"

type NotEmptyValidation struct {
}

func (n NotEmptyValidation) Tag() string {
	//TODO implement me
	return "notempty"
}

func (n NotEmptyValidation) Valid(fl validator.FieldLevel) bool {
	return fl.Field().Len() > 0
}

func (n NotEmptyValidation) I18n() (section string, key string) {
	//TODO implement me
	return "validation.default", "notempty"
}
