package config

import (
	"go/riada/routes"
	"go/riada/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Server() {
	port := utils.Dotenv("SERVER_PORT")
	if port == "" {
		port = "3003"
	}
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./client/dist")

	routes.AllRoutes(app)

	// database.Migrate()

	app.Listen(":" + port)

}
