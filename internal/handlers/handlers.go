package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/hackertron/go-chess/internal/db"
	"github.com/hackertron/go-chess/internal/middlewares"
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
	authenticated, err := db.AuthenticateUser(dbs, username, password)
	if err != nil {
		return err
	}
	defer db.CloseDB(dbs)
	if authenticated {
		// return c.JSON(http.StatusOK, map[string]interface{}{
		// 	"user":     user,
		// 	"userInfo": userInfo,
		// })
		cookie := new(http.Cookie)
		cookie.Name = "sessionID"
		cookie.Value = uuid.New().String()
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(24 * time.Hour)
		c.SetCookie(cookie)
		userSession := c.Get("sessions").(*middlewares.Sessions)
		userSession.AddSession(cookie.Value, username)
		c.Set("sessions", &userSession)
		// redirect to dashboard page
		return c.Redirect(http.StatusFound, "/user/dashboard")
		// return render(c, views.Dashboard(showBase))
	}
	return echo.ErrUnauthorized

}

func LogoutUser(c echo.Context) error {
	// Clear the session cookie to log the user out
	cookie := new(http.Cookie)
	cookie.Name = "sessionID"
	cookie.Value = ""
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	// Redirect to the login page after logout
	return render(c, views.Login(false))
}

func DashboardPage(c echo.Context) error {
	cookie, err := c.Cookie("sessionID")
	if err != nil {
		log.Fatal("error in getting cookie : ", err)
	}
	session := c.Get("sessions").(*middlewares.Sessions)
	username := session.GetSession(cookie.Value)
	dbs, err := db.ConnectToDB("chess.db")
	if err != nil {
		return err
	}
	defer db.CloseDB(dbs)
	user, userInfo, err := db.GetUserFromUsername(dbs, username)
	if err != nil {
		return err
	}
	fmt.Println(user)
	fmt.Println(userInfo)
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"user":     user,
	// 	"userInfo": userInfo,
	// })
	return render(c, views.Dashboard(false, user, userInfo))
}
