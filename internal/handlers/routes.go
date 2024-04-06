package handlers

import (
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	group := app.Group("/user")
	group.POST("/create", CreateUser)
	group.GET("/create", CreateUserPage)
	group.GET("/login", LoginUserPage)
	group.POST("/login", LoginUser)
}
