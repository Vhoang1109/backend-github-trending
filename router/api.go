package router

import (
	"backend-github-trending/handler"
	middleware "backend-github-trending/middleware"

	"github.com/labstack/echo"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handler.UserHandler
}

func (api *API) SetupRouter() {
	api.Echo.POST("/user/sign-in", api.UserHandler.HandleSignIn)
	api.Echo.POST("/user/sign-up", api.UserHandler.HandleSignUp)

	// Using a different route for Profile and applying the JWT middleware correctly
	api.Echo.GET("/user/profile", api.UserHandler.Profile, middleware.JWTMiddleware())
}
