package authentication

import (
	"fmt"
	"go/riada/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	JwtSecretKey     = []byte(utils.Dotenv("SECRET_KEY"))
	jwtSigningMethod = jwt.SigningMethodHS256.Name
)

func NewToken(userId string, username string, remember bool) (string, error) {
	Expires := jwt.NewNumericDate(time.Now())
	if remember {
		Expires = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	} else {
		Expires = jwt.NewNumericDate(time.Now().Add(1 * time.Hour))
	}

	claims := jwt.RegisteredClaims{

		ExpiresAt: Expires,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    userId,
		Subject:   username,
		ID:        userId,
		Audience:  []string{"*"},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecretKey)

}

func validateSignedMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	}
	return JwtSecretKey, nil
}

func ParseToken(tokenString string) (*jwt.RegisteredClaims, error) {
	claims := new(jwt.RegisteredClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, validateSignedMethod)
	if err != nil {
		return nil, err
	}
	var ok bool
	claims, ok = token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return nil, utils.ErrInvalidAuthToken
	}
	return claims, nil
}

// func validateSignedMethod(token *jwt.Token)(interface{},error){
// 	if _, ok := jwt.Parse(token); !ok{
// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 	}
// 	return JwtSecretKey,nil
// }
