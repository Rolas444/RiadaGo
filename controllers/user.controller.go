package controllers

import (
	"errors"
	"go/riada/connection"
	"go/riada/models"
	"go/riada/utils/authentication"
	"go/riada/utils/sessionjwt"

	"github.com/go-passwd/validator"
	"github.com/gofiber/fiber/v2"
)

type UserLogin struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Remember bool   `json:"remember"`
}

type UserResponse struct {
	UserName string
	Role     string
	Id       uint
	login    bool
}

func LoginUser(c *fiber.Ctx) error {
	user := new(UserLogin)
	var userbd models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(&err)
	}
	// fmt.Println(user)
	db := connection.Db()
	db.Where(&models.User{UserName: user.UserName}).First(&userbd)
	err := authentication.VerifyPass(userbd.Password, user.Password)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	session, err := sessionjwt.SignedLoginToken(&userbd)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(session)
}

func CreateUser(c *fiber.Ctx) error {
	e := errors.New("Password invalid : must have more than 5 letters and special characters")
	nuser := new(models.User)
	if err := c.BodyParser(&nuser); err != nil {
		return c.Status(400).JSON(&err)
	}
	passValidator := validator.New(validator.MinLength(6, e), validator.CommonPassword(e))
	err := passValidator.Validate(nuser.Password)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	npass, err := authentication.EncryptPass(nuser.Password)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	db := connection.Db()

	nuser.Password = npass
	db.Create(&nuser)
	// return c.Status(200).JSON(nuser)
	ruser := new(UserResponse)
	ruser.UserName = nuser.UserName
	ruser.Role = nuser.Role
	ruser.Id = nuser.ID
	return c.Status(200).JSON(ruser)
}
