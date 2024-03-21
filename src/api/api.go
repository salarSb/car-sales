package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/salarSb/car-sales/api/middlewares"
	"github.com/salarSb/car-sales/api/routers"
	"github.com/salarSb/car-sales/api/validations"
	"github.com/salarSb/car-sales/config"
	"log"
)

func InitServer() {
	cfg := config.GetConfig()
	r := gin.New()
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
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequestMiddleware())
	api := r.Group("api")
	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		testRouter := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(testRouter)
	}
	v2 := api.Group("/v2/")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		log.Fatal("error on running router")
		return
	}
}
