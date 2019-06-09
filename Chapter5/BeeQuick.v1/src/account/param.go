package account

import (
	"fmt"
	"unicode"

	"gopkg.in/go-playground/validator.v9"
)

type RegisterParam struct {
	Phone    string `form:"phone" json:"phone" validate:"required,len=11"`
	Password string `form:"password" json:"password"`
}

// 参数校验
func (param RegisterParam) suitable() (bool, error) {
	if param.Password == "" || len(param.Phone) != 11 {
		return false, fmt.Errorf("password should not be nil or the length of phone is not 11")
	}
	if unicode.IsNumber(rune(param.Password[0])) {
		return false, fmt.Errorf("password should start with number")
	}
	return true, nil
}

// 参数校验使用 Tag 检查
func registerValidation(sl validator.StructLevel) {
	param := sl.Current().Interface().(RegisterParam)
	if param.Phone == "" && param.Password == "" {
		sl.ReportError(param.Password, "Password", "password", "password", param.Password)
		sl.ReportError(param.Phone, "Phone", "phone", "phone", param.Phone)
	}
}

func (param RegisterParam) Valid() *validator.Validate {
	validate := validator.New()
	validate.RegisterStructValidation(registerValidation, RegisterParam{})
	return validate
}
