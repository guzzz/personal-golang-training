package autentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.JwtSecret))
}

// ValidateToken será responsável por validar o Token
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, erro := jwt.Parse(tokenString, getVerificationKey)
	if erro != nil {
		return errors.New("Usuário não autenticado.")
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido!")
}

// ExtractUserIdFromToken será responsável por extrair o userId do Token
func ExtractUserIdFromToken(r *http.Request) (uint64, error) {
	tokenString := extractToken(r)
	token, _ := jwt.Parse(tokenString, getVerificationKey)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenUserIdString := fmt.Sprintf("%.0f", claims["userId"])
		tokenUserId, erro := strconv.ParseUint(tokenUserIdString, 10, 64)
		if erro != nil {
			return 0, erro
		}
		return tokenUserId, nil
	}
	return 0, errors.New("Token inválido!")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.JwtSecret, nil
}
