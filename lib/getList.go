package lib

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Get is getList
func Get(c echo.Context) error {
	url := "https://280blocker.net/files/280blocker_adblock_" + time.Now().Format("200601") + ".txt"
	return c.Redirect(http.StatusFound, url)
}
