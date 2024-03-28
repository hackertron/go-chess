package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(app *echo.Echo, h *UserHandler) {
	group := app.Group("/user")
	group.GET("", h.HandlerShowUsers)
	group.GET("/details/:id", h.HandlerShowUserById)
	// add registration and login route
	group.GET("/register", h.HandlerRegisterUser)
	group.POST("/register", h.HandlerRegisterUser)
	//group.POST("/login", h.HandlerLoginUser)

}
