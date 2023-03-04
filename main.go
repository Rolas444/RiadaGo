package main

import (
	"go/riada/config"
	"go/riada/database"
)

func main() {
	config.Server()
	database.Migrate()

}
