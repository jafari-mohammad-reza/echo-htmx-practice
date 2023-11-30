package main

import (
	"github.com/jafari-mohammad-reza/echo-htmx-practice/pkg"
	"github.com/jafari-mohammad-reza/echo-htmx-practice/pkg/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"net/http"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))
	pkg.NewTemplateRenderer(e, "views/*.html")
	err := db.SeedInit()
	if err != nil {
		panic(err)
	}

	e.GET("/hello", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index.html", map[string]string{
			"Message": "Hello world.",
		})
	})

	e.Logger.Fatal(e.Start(":4040"))
}
