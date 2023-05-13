package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	URI                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequestAutentication bool
}

// Settings configura todas as rotas dentro do Router
func Settings(r *mux.Router) *mux.Router {
	routes := usersRoutes
	routes = append(routes, loginRoutes...)
	routes = append(routes, publicationsRoutes...)

	for _, route := range routes {

		if route.RequestAutentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Autenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
