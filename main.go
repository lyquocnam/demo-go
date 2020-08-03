package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		req := c.Request()
		format := `
    <code>
      Protocol: %s<br>
      Host: %s<br>
      Remote Address: %s<br>
      Method: %s<br>
      Path: %s<br>
    </code>
  `
		return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL.Path))
	})

	e.Logger.Fatal(e.StartTLS(":3001",
		"/etc/letsencrypt/live/daotuoi.com/fullchain.pem",
		"/etc/letsencrypt/live/daotuoi.com/privkey.pem",
		))
}
