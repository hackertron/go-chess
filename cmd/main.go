package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/hackertron/go-chess/internal/db"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type db_sql struct {
	db *sql.DB
}

func main() {
	fmt.Println("Go-Chess go brrr 🚀🚀")

	e := echo.New()
	db_sql, err := db.ConnectToDB("chess.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseDB(db_sql)
	db.CreateMigrationsTable(db_sql)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	group := e.Group("/user")
	group.POST("/create", func(c echo.Context) error {
		var users db.Users
		if err := c.Bind(&users); err != nil {
			return err
		}
		user, userInfo, err := db.CreateUser(db_sql, users)
		if err != nil {
			return err
		}
		fmt.Println(user)
		fmt.Println(userInfo)
		return c.JSON(http.StatusOK, user)
	})

	e.Logger.Fatal(e.Start("localhost:8080"))

}
