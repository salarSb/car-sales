package main

import (
	"github.com/salarSb/car-sales/api"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/cache"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.StartUp, err.Error(), nil)
		return
	}
	defer cache.CloseRedis()
	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.StartUp, err.Error(), nil)
		return
	}
	defer db.CloseDb()
	api.InitServer(cfg)
}
