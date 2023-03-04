package database

import (
	"go/riada/config"
	"go/riada/models"
	"log"
)

func Migrate() {
	log.Printf("Migrate: Start ...")
	db := config.Db()
	db.AutoMigrate(
		&models.Ministry{},
		&models.TypeDoc{},
		&models.Sex{},
		&models.Fellow{},
		&models.Address{},
		&models.Comunity{},
	)
	log.Printf("Migrate: Finished")

}
