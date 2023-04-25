package routes

import (
	"go/riada/controllers"

	"github.com/gofiber/fiber/v2"
)

func MinistryRoutes(router *fiber.App) {

	ministry := router.Group("/ministry")
	ministry.Post("/", controllers.CreateMinistry)
	ministry.Get("/", controllers.GetMinistries)
}
