package main

import (
	"net/http"

	"github.com/jafari-mohammad-reza/echo-htmx-practice/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Debug = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	pkg.NewTemplateRenderer(e, "views/*.html")
	go func() {
		err := pkg.InitDatabase()
		if err != nil {
			panic(err)
		}
	}()
	setupRoutes(e)
	e.Logger.Fatal(e.Start(":4040"))
}

func setupRoutes(e *echo.Echo) {
	e.GET("/", func(e echo.Context) error {
		return e.Render(http.StatusOK, "index.html", map[string]string{
			"Title": "Htmx Echo practice.",
		})
	})
	bl := e.Group("/blog")
	bl.GET("", pkg.GetBlogs)
	bl.GET(":id", pkg.GetBlog)
	bl.POST("", pkg.CreateBlog)
	bl.DELETE("", pkg.DeleteBlog)
}
