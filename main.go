package main

import (
	"go/riada/config"
	"go/riada/connection"
	"go/riada/database"
	"go/riada/seeds"
)

func main() {
	db := connection.Db()
	database.Migrate(db)
	seeds.SetSexes(db)
	seeds.SetTypeDocs(db)
	config.Server()
}
