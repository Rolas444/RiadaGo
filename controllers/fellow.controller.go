package controllers

import (
	"go/riada/connection"
	"go/riada/models"

	"github.com/gofiber/fiber/v2"
)

// type Fellow struct{
// 	ID uint `json:"id"`
// 	Name string `json:""`
// }

// func CreateResponseFellow( name, ) models.Fellow {
// 	return models.Fellow{}
// }

func GetFellows(c *fiber.Ctx) error {
	db := connection.Db()
	fellows := []models.Fellow{}

	db.Find(&fellows)

	return c.JSON(fellows)

}
