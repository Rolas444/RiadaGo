package authentication

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPass(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "no se pudo encriptar", err
	}

	return string(hashed), nil
}

func VerifyPass(hashed, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
}

func VerifyToken(c *fiber.Ctx) (string, error) {
	auth := c.Request().Header.Peek("Authorization")
	claims, err := ParseToken(string(auth))
	if err != nil {
		return err.Error(), err
	}
	return claims.Subject, nil
}
