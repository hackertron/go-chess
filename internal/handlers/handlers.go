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
	fmt.Println("users we got  : ", users)
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
	var showBase = false
	return render(c, views.Login(showBase))
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

func CreateUserPage(c echo.Context) error {
	return render(c, views.Register())
}

func LoginUserPage(c echo.Context) error {
	var showBase = true
	return render(c, views.Login(showBase))
}

func LoginUser(c echo.Context) error {
	dbs, err := db.ConnectToDB("chess.db")
	if err != nil {
		return err
	}
	username := c.FormValue("username")
	password := c.FormValue("password")
	user, userInfo, err := db.AuthenticateUser(dbs, username, password)
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
	var showBase = false
	return render(c, views.Login(showBase))
}
