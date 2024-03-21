package main

import (
	"github.com/salarSb/car-sales/api"
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/cache"
	"github.com/salarSb/car-sales/data/db"
	"log"
)

func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.CloseDb()
	api.InitServer(cfg)
}
