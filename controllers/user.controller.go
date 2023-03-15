package controllers

import (
	"errors"
	"fmt"
	"go/riada/connection"
	"go/riada/models"
	"go/riada/utils/authentication"
	"net/http"
	"strconv"

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
	var userdb models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(&err)
	}
	// fmt.Println(user)
	db := connection.Db()
	db.Where(&models.User{UserName: user.UserName}).First(&userdb)
	err := authentication.VerifyPass(userdb.Password, user.Password)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	// session, err := sessionjwt.SignedLoginToken(&userbd)
	// if err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }

	token, err := authentication.NewToken(strconv.Itoa(int(userdb.ID)), userdb.UserName, user.Remember)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	ur := new(UserResponse)
	ur.Id = userdb.ID
	ur.UserName = userdb.UserName
	ur.Role = userdb.Role
	ur.login = true

	return c.Status(http.StatusOK).
		JSON(fiber.Map{
			"user":  ur,
			"token": fmt.Sprintf(token),
		})
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
