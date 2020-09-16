package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World2")
}

// Handler is router
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/", hello)
	e.GET("/api/2", hello2)

	e.ServeHTTP(w, r)
}
