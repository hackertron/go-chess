package handlers

import (
	"github.com/hackertron/go-chess/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(app *echo.Echo) {
	group := app.Group("/user")
	group.POST("/create", CreateUser)
	group.GET("/create", CreateUserPage)
	group.GET("/login", LoginUserPage)
	group.POST("/login", LoginUser)
	group.GET("/dashboard", DashboardPage, middlewares.AuthenticatedUser)

	app.GET("/ws", WebSocketHandler)
	app.GET("/wsc", WscPage)
}
