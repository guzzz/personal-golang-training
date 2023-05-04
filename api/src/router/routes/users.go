package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                  "/users",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.ListUsers,
		RequestAutentication: false,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.RetrieveUser,
		RequestAutentication: false,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodPut,
		Function:             controllers.UpdateUser,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/{userId}/follow",
		Method:               http.MethodPost,
		Function:             controllers.FollowUser,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/{userId}/unfollow",
		Method:               http.MethodPost,
		Function:             controllers.UnfollowUser,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/{userId}/followers",
		Method:               http.MethodGet,
		Function:             controllers.GetFollowers,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/{userId}/following",
		Method:               http.MethodGet,
		Function:             controllers.GetFollowing,
		RequestAutentication: true,
	},
	{
		URI:                  "/users/update-password",
		Method:               http.MethodPost,
		Function:             controllers.UpdatePassword,
		RequestAutentication: true,
	},
}
