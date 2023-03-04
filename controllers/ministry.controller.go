package controllers

import (
	"go/riada/connection"
	"go/riada/models"

	"github.com/gofiber/fiber/v2"
)

func CreateMinistry(c *fiber.Ctx) error {
	db := connection.Db()
	var ministry models.Ministry
	if err := c.BodyParser(&ministry); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	db.Create(&ministry)
	return c.Status(200).JSON(ministry)
}
