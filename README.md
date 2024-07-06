# github.com/kordar/govalidator

对`validator/v10`简单的抽象，自定义`tag`，自定义验证逻辑，自定义模版，多语言。适用于模版配置，多语言等场景。

```go
type IValidation interface {
	Tag() string
	Valid(fl validator.FieldLevel) bool
	DefaultTpl() (tpl string, override bool)
	Tpl() (section string, key string)
	I18n(fe validator.FieldError, locale string) []string
}
```

## 组件初始化

```go
validate := validator.New()
govalidator.LoadValidate(validate)
```

- gin框架参考使用
```go
import (
    "github.com/gin-gonic/gin/binding"
    "github.com/go-playground/validator/v10"
    "github.com/kordar/govalidator"
)

validate := binding.Validator.Engine().(*validator.Validate)
govalidator.LoadValidate(validate)
```


## 配合多语言实现

- [国际化组件](github.com/go-playground/universal-translator "国际化")
- [gotrans集成国际化组件](github.com/kordar/gotrans)

将实现对象注入国际化组件

```go
import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

validate := validator.New()
var u ut.Translator // 国际化
var v IValidation  // .....
var tag = v.Tag()  
validate.RegisterTranslation(tag, u, func(ut ut.Translator) error {
	return ut.Add(tag, v.Tpl(), override)
}, func(ut ut.Translator, fe validator.FieldError) string {
	tt := v.I18n()
    value := ut.T(tag, tt...)
	return value
})
```

