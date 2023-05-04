package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Connect abre a conexão com o banco
func Connect() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringDBConection)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
