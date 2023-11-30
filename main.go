package main

import (
	"github.com/jafari-mohammad-reza/echo-htmx-practice/src/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
)

type Msg struct {
	Message string
}

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	pkg.NewTemplateRenderer(e, "views/*.html")
	e.GET("/hello", func(e echo.Context) error {
		msg := Msg{Message: "Hello world"}
		return e.Render(http.StatusOK, "index.html", msg)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
