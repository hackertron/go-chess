package middlewares

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Sessions struct {
	Sessions map[string]string
}

func NewSessions() *Sessions {
	return &Sessions{
		Sessions: make(map[string]string),
	}
}

func (s *Sessions) AddSession(key string, value string) {
	s.Sessions[key] = value
}

func (s *Sessions) GetSession(key string) string {
	return s.Sessions[key]
}

func (s *Sessions) DeleteSession(key string) {
	delete(s.Sessions, key)
}

func AuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		fmt.Println("cookied : ", cookie)
		if err != nil || cookie.Value == "" {
			// Redirect to login page if session cookie doesn't exist
			return c.Redirect(http.StatusFound, "/user/login")
		}
		// check with memory sessions
		sessions := c.Get("sessions").(*Sessions)
		if _, ok := sessions.Sessions[cookie.Value]; !ok {
			// Redirect to login page if session cookie doesn't exist in memory sessions
			return c.Redirect(http.StatusFound, "/user/login")
		}
		// Continue to the next handler if user is logged in
		return next(c)
	}
}
