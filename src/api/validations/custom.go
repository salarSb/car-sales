package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/salarSb/car-sales/common"
	"log"
	"regexp"
)

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}
	res, err := regexp.MatchString(`^09(0[0-9]|1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}`, value)
	if err != nil {
		log.Print(err.Error())
	}
	return res
}

func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}
	return common.CheckPassword(value)
}
