package handlers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/hackertron/go-chess/internal/db"
	"github.com/hackertron/go-chess/internal/views"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, cmp templ.Component) error {
	return cmp.Render(ctx.Request().Context(), ctx.Response())
}

func CreateUser(c echo.Context) error {
	var users db.Users
	dbs, err := db.ConnectToDB("chess.db")
	if err != nil {
		return err
	}
	if err := c.Bind(&users); err != nil {
		return err
	}
	user, userInfo, err := db.CreateUser(dbs, users)
	if err != nil {
		return err
	}
	defer db.CloseDB(dbs)
	fmt.Println(user)
	fmt.Println(userInfo)
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"user":     user,
	// 	"userInfo": userInfo,
	// })
	return render(c, views.Home(user, userInfo))
}

func GetUser(c echo.Context) error {
	dbs, err := db.ConnectToDB("chess.db")
	if err != nil {
		return err
	}
	defer db.CloseDB(dbs)
	user, userInfo, err := db.GetUser(dbs)
	if err != nil {
		return err
	}
	return render(c, views.Details(user, userInfo))
}
