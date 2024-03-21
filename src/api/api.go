package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/salarSb/car-sales/api/middlewares"
	"github.com/salarSb/car-sales/api/validations"
	"github.com/salarSb/car-sales/config"
	"log"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterValidators()
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequestMiddleware())
	RegisterRoutes(r, cfg)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal("error on running router")
		return
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

}

func RegisterValidators() {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			log.Fatal("Error on registering custom validations")
			return
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			log.Fatal("Error on registering custom validations")
			return
		}
	}
}
