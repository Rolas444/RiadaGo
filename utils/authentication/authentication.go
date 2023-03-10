package authentication

import "golang.org/x/crypto/bcrypt"

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
