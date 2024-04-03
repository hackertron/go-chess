package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/hackertron/go-chess/internal/db"
	"github.com/hackertron/go-chess/internal/handlers"
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
	db.CreateMigrationsTable(db_sql)
	db.CloseDB(db_sql)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	handlers.SetupRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))

}
