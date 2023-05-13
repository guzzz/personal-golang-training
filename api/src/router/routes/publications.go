package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationsRoutes = []Route{
	{
		URI:                  "/publications",
		Method:               http.MethodPost,
		Function:             controllers.CreatePublication,
		RequestAutentication: false,
	},
	{
		URI:                  "/publications",
		Method:               http.MethodGet,
		Function:             controllers.ListPublications,
		RequestAutentication: false,
	},
	{
		URI:                  "/my-publications",
		Method:               http.MethodGet,
		Function:             controllers.ListMyPublications,
		RequestAutentication: false,
	},
	{
		URI:                  "/publications/{publicationId}",
		Method:               http.MethodGet,
		Function:             controllers.RetrievePublication,
		RequestAutentication: false,
	},
	{
		URI:                  "/publications/{publicationId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdatePublication,
		RequestAutentication: true,
	},
	{
		URI:                  "/publications/{publicationId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeletePublication,
		RequestAutentication: true,
	},
	{
		URI:                  "/publications/{publicationId}/like",
		Method:               http.MethodPost,
		Function:             controllers.Like,
		RequestAutentication: true,
	},
	{
		URI:                  "/publications/{publicationId}/remove-like",
		Method:               http.MethodPost,
		Function:             controllers.RemoveLike,
		RequestAutentication: true,
	},
}
