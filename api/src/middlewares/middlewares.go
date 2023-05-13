package middlewares

import (
	"api/src/autentication"
	"api/src/responses"
	"log"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n%s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Autenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := autentication.ValidateToken(r); erro != nil {
			responses.Error(w, http.StatusUnauthorized, erro)
			return
		}
		next(w, r)
	}
}
