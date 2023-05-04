package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON retorna uma resposta JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}

// Error retorna um erro para a requisição
func Error(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
