package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func getUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, "done")
}
