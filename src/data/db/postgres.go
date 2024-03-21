package db

import (
	"fmt"
	"github.com/salarSb/car-sales/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	cnn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
		cfg.Postgres.SSLMode,
	)
	dbClient, err := gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != nil {
		return err
	}
	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConnections)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConnections)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnectionMaxLifetime * time.Minute)
	log.Println("db connection established")
	return nil
}

func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	cnn, _ := dbClient.DB()
	err := cnn.Close()
	if err != nil {
		log.Fatal("error on closing db connection")
		return
	}
}
