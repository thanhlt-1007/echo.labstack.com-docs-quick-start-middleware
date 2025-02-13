package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func getUsersHandler(context echo.Context) error {
	return context.String(
		http.StatusOK,
		"OK",
	)
}

func trackMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		println("request to /users")
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/users", getUsersHandler, trackMiddleware)
	e.Logger.Fatal(e.Start(":1323"))
}
