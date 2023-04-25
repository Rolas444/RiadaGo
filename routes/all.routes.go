package routes

import "github.com/gofiber/fiber/v2"

func AllRoutes(router *fiber.App) {
	// routes := router.Group("/api")
	MinistryRoutes(router)
	UserRoutes(router)

}
