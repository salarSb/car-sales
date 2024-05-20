package main

import (
	"github.com/salarSb/car-sales/api"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/cache"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/db/migrations"
)

func main() {
	cfg := config.GetConfig()
	cache.GetRedis()
	defer cache.CloseRedis()
	db.GetDb()
	migrations.Up1()
	defer db.CloseDb()
	api.InitServer(cfg)
}
