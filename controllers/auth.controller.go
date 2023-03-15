package controllers

import (
	"go/riada/connection"
	"go/riada/models"
	"go/riada/utils"
	"go/riada/utils/authentication"

	"github.com/gofiber/fiber/v2"
)

func Auth(c *fiber.Ctx, roles [3]string) (bool, error) {
	subject, err := authentication.VerifyToken(c)
	if err != nil {
		return false, err
	}
	var user models.User
	db := connection.Db()
	db.Where(&models.User{UserName: subject}).First(&user)
	if utils.ItemExists(roles, user.Role) {
		return true, nil
	}

	return false, nil
}
