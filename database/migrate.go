package database

import (
	"go/riada/models"
	"log"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	log.Printf("Migrate: Start ...")
	// db := connection.Db()
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
