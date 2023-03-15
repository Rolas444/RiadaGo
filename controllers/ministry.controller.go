package controllers

import (
	"errors"
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

func GetMinistries(c *fiber.Ctx) error {
	// var roles []string
	// roles[0]="user"
	// roles[1]="admin"
	roles := [3]string{"user", "other"}
	auth, err := Auth(c, roles)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{"Error": "Internal Error"})
	}
	if auth == false {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{"message": "Unauthorized user", "status": "401"})
	}
	var ministries []models.Ministry
	db := connection.Db()
	db.Where(&models.Ministry{Status: 1}).Find(&ministries)
	if len(ministries) == 0 {
		return c.Status(400).JSON(errors.New("No found items"))
	}

	return c.Status(200).JSON(ministries)

}
