package lib

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// Get is getList
func Get(c echo.Context) error {
	urls := []string{"https://280blocker.net/files/280blocker_adblock_" + time.Now().Format("200601") + ".txt"}
	return c.String(http.StatusOK, getUrls(urls))
}

func getUrls(urls []string) string {
	var lists = ""
	for _, url := range urls {
		fmt.Println(url)
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
