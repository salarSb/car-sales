package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/salarSb/car-sales/api/helper"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/constants"
	"github.com/salarSb/car-sales/pkg/service_errors"
	"github.com/salarSb/car-sales/services"
	"net/http"
	"strings"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	var err error
	var tokenService = services.NewTokenService(cfg)
	return func(context *gin.Context) {
		claimMap := map[string]interface{}{}
		auth := context.GetHeader(constants.AuthorizationHeaderKey)
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			token := strings.Split(auth, " ")
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				helper.GenerateBaseResponseWithError(nil, false, helper.AuthError, err),
			)
			return
		}
		context.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		context.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		context.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		context.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		context.Set(constants.EmailKey, claimMap[constants.EmailKey])
		context.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		context.Set(constants.RolesKey, claimMap[constants.RolesKey])
		context.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])
		context.Next()
	}
}
