package main

import "net/http"

type Msg struct {
	Message string
}

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(
		rate.Limit(20),
	)))

	template.NewTemplateRenderer(e, "public/*.html")
	e.GET("/hello", func(e echo.Context) error {
		msg := Msg{Message: "Hello world"}
		return c.Render(http.StatusOK, "index", msg)
	})

	e.Logger.Fatal(e.Start(":4040"))
}
