package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/salarSb/car-sales/api/middlewares"
	"github.com/salarSb/car-sales/api/routers"
	"github.com/salarSb/car-sales/api/validations"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/docs"
	"github.com/salarSb/car-sales/pkg/logging"
	"github.com/salarSb/car-sales/pkg/metrics"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	logger := logging.NewLogger(cfg)
	gin.SetMode(cfg.Server.RunMode)
	r := gin.New()
	RegisterValidators(logger)
	RegisterPrometheus(logger)
	r.Use(middlewares.DefaultStructuredLogger(cfg))
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.Prometheus())
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler))
	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.InternalPort))
	if err != nil {
		logger.Fatal(logging.Internal, logging.Api, "error on running router", nil)
		return
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		//User
		users := v1.Group("/users")

		//Base
		countries := v1.Group(
			"/countries",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		cities := v1.Group(
			"/cities",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		files := v1.Group(
			"/files",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		colors := v1.Group(
			"/colors",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		years := v1.Group(
			"/years",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)

		//Properties
		propertyCategories := v1.Group(
			"/property-categories",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		properties := v1.Group(
			"/properties",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)

		//car
		carTypes := v1.Group(
			"/car-types",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		gearboxes := v1.Group(
			"/gearboxes",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		companies := v1.Group(
			"/companies",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModels := v1.Group(
			"/car-models",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelColors := v1.Group(
			"/car-model-colors",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelYears := v1.Group(
			"/car-model-years",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelPriceHistories := v1.Group(
			"/car-model-price-histories",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelFiles := v1.Group(
			"/car-model-files",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelProperties := v1.Group(
			"/car-model-properties",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelCommentsAdmin := v1.Group(
			"/car-model-comments",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin"}),
		)
		carModelCommentsUser := v1.Group(
			"/car-model-comments",
			middlewares.Authentication(cfg),
			middlewares.Authorization([]string{"admin", "default"}),
		)

		//User
		routers.User(users, cfg)

		//Base
		routers.Country(countries, cfg)
		routers.City(cities, cfg)
		routers.File(files, cfg)
		routers.Color(colors, cfg)
		routers.Year(years, cfg)

		//Properties
		routers.PropertyCategory(propertyCategories, cfg)
		routers.Property(properties, cfg)

		//car
		routers.CarType(carTypes, cfg)
		routers.Gearbox(gearboxes, cfg)
		routers.Company(companies, cfg)
		routers.CarModel(carModels, cfg)
		routers.CarModelColor(carModelColors, cfg)
		routers.CarModelYear(carModelYears, cfg)
		routers.CarModelPriceHistory(carModelPriceHistories, cfg)
		routers.CarModelFile(carModelFiles, cfg)
		routers.CarModelProperty(carModelProperties, cfg)
		routers.CarModelCommentAdmin(carModelCommentsAdmin, cfg)
		routers.CarModelCommentUser(carModelCommentsUser, cfg)

		r.Static("/static", "./uploads")

		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
}

func RegisterValidators(logger logging.Logger) {
	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := val.RegisterValidation("mobile", validations.IranianMobileNumberValidator, true)
		if err != nil {
			logger.Fatal(
				logging.Validation,
				logging.MobileValidation,
				"Error on registering custom mobile validation",
				nil,
			)
			return
		}
		err = val.RegisterValidation("password", validations.PasswordValidator, true)
		if err != nil {
			logger.Fatal(
				logging.Validation,
				logging.PasswordValidation,
				"Error on registering custom password validation",
				nil,
			)
			return
		}
	}
}

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api for car sale project"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func RegisterPrometheus(logger logging.Logger) {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.StartUp, err.Error(), nil)
	}
	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.StartUp, err.Error(), nil)
	}
}
