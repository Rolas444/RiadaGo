package routes

import (
	"go/riada/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router *fiber.App) {
	user := router.Group("/user")
	user.Post("/login", controllers.LoginUser)
	user.Post("/register", controllers.CreateUser)

}
