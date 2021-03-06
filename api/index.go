package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sh4869221b/go-app/lib"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func hello2(c echo.Context) error {
	return c.String(http.StatusOK, "https://280blocker.net/files/280blocker_adblock_"+time.Now().Format("200601")+".txt")
}

// Handler is router
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/", hello)
	e.GET("/api/2", hello2)
	e.GET("/api/list.txt", lib.Get)

	e.ServeHTTP(w, r)
}
