package main

import (
	"go/riada/config"
	"go/riada/database"
)

func main() {

	database.Migrate()
	config.Server()
}
