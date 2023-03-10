package sessionjwt

import (
	"go/riada/models"
	"go/riada/utils"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	UseName string
	Role    string
}

type Claim struct {
	User `json:"user"`
	jwt.Claims
}

type ResponseToken struct {
	Token string `json:"token"`
}

func SignedLoginToken(u *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   u.ID,
		"username": u.UserName,
		"role":     u.Role,
		"login":    true,
	})

	return token.SignedString([]byte(utils.Dotenv("SECRET_KEY")))
}
