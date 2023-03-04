package config

import (
	"fmt"
	"go/riada/utils"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Db() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Lima",
		utils.Dotenv("DB_HOST"),
		utils.Dotenv("DB_USER"),
		utils.Dotenv("DB_PASSWORD"),
		utils.Dotenv("DB_NAME"),
	)
	log.Println(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	return db
}
