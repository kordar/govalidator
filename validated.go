package govalidator

import (
	"github.com/go-playground/validator/v10"
)

var _validate *validator.Validate

type IValidation interface {
	Tag() string
	Valid(fl validator.FieldLevel) bool
	I18n() (section string, key string)
}

func LoadValidate(valid *validator.Validate) {
	_validate = valid
}

func GetValidate() *validator.Validate {
	return _validate
}

func AddValidation(validation IValidation) {
	_ = _validate.RegisterValidation(validation.Tag(), validation.Valid)
}
