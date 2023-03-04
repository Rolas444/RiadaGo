package database

import (
	"go/riada/connection"
	"go/riada/models"
	"log"
)

func Migrate() {
	log.Printf("Migrate: Start ...")
	db := connection.Db()
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
