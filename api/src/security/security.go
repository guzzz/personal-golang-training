package security

import (
	"api/src/config"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) (string, error) {
	cost := config.BcryptCost

	encrypted_pwd_in_bytes, erro := bcrypt.GenerateFromPassword([]byte(password), cost)
	if erro != nil {
		return "", erro
	}
	encrypted_pwd := string(encrypted_pwd_in_bytes[:])
	return encrypted_pwd, nil
}

func CompareHash(password, hashedPassword string) error {
	if erro := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)); erro != nil {
		return errors.New("Falha na autenticação do usuário.")
	} else {
		return nil
	}
}
