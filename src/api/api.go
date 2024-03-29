package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/salarSb/car-sales/api/middlewares"
	"github.com/salarSb/car-sales/api/validations"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/docs"
	"github.com/salarSb/car-sales/pkg/logging"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	logger := logging.NewLogger(cfg)
	r := gin.New()
	RegisterValidators(logger)
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middlewares.LimitByRequestMiddleware())
	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
	if err != nil {
		logger.Fatal(logging.Internal, logging.Api, "error on running router", nil)
		return
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

}

func RegisterValidators(logger logging.Logger) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Fatal(logging.Validation, logging.MobileValidation, "Error on registering custom mobile validation", nil)
			return
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			logger.Fatal(logging.Validation, logging.PasswordValidation, "Error on registering custom password validation", nil)
			return
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api for car sale project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
