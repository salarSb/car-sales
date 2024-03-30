package helper

import (
	"github.com/salarSb/car-sales/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	//Otp
	service_errors.OtpExists:   http.StatusConflict,
	service_errors.OtpUsed:     http.StatusConflict,
	service_errors.OtpNotValid: http.StatusBadRequest,
}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
