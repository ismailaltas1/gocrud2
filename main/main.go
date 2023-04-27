package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/users", getUsers)

	e.Start(":8080")

}
