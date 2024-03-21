package main

import (
	"github.com/salarSb/car-sales/api"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
