package migrations

import (
	"github.com/salarSb/car-sales/config"
	"github.com/salarSb/car-sales/data/db"
	"github.com/salarSb/car-sales/data/models"
	"github.com/salarSb/car-sales/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := db.GetDb()
	var tables []interface{}
	country := models.Country{}
	city := models.City{}
	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}
	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}
	err := database.Migrator().CreateTable(tables...)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Migration, "cannot create tables through migration", nil)
		return
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func Down1() {

}
