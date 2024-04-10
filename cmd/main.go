package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hackertron/go-chess/internal/db"
	"github.com/hackertron/go-chess/internal/handlers"
	"github.com/hackertron/go-chess/internal/middlewares"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

//	type db_sql struct {
//		db *sql.DB
//	}
func main() {
	fmt.Println("Go-Chess go brrr 🚀🚀")

	e := echo.New()
	// static files
	e.Static("/static", "assets")
	db_sql, err := db.ConnectToDB("chess.db")
	if err != nil {
		log.Fatal(err)
	}
	db.CreateMigrationsTable(db_sql)
	db.CloseDB(db_sql)

	sessions := middlewares.NewSessions()
	e.Pre(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("sessions", sessions)
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	handlers.SetupRoutes(e)

	e.Logger.Fatal(e.Start("localhost:8080"))

}
