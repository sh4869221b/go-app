package handler

import (
	"io/ioutil"
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

func get(c echo.Context) error {
	urls := []string{"https://280blocker.net/files/280blocker_adblock.txt", "https://raw.githubusercontent.com/nanj-adguard/nanj-filter/master/nanj-filter.txt", "https://blog-imgs-116-origin.fc2.com/b/t/o/btonews/5ch_matome_filter.txt", "https://raw.githubusercontent.com/tofukko/filter/master/Adblock_Plus_list.txt"}
	return c.String(http.StatusOK, getUrls(urls))
}

func getUrls(urls []string) string {
	var lists = ""
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		lists += string(body)
	}
	return lists
}

// Handler is router
func Handler(w http.ResponseWriter, r *http.Request) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/api/", hello)
	e.GET("/api/2", hello2)
	e.GET("/api/list.txt", get)

	e.ServeHTTP(w, r)
}
