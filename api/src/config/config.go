package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringDBConection é a string de conexão com o MySQL
	StringDBConection = ""

	// Port é a porta onde o DB vai rodar
	Port = 0

	// BcryptCost é o valor do custo da parte de criptografia
	BcryptCost = 0

	// JwtSecret é o valor utilizado como chave ao se tratar de JWTs
	JwtSecret []byte
)

// Load vai incializar variaveis de ambiente
func Load() {
	var erro error = godotenv.Load()
	if erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringDBConection = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
	)

	BcryptCost, erro = strconv.Atoi(os.Getenv("BCRYPT_COST"))
	if erro != nil {
		BcryptCost = 10
	}

	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
