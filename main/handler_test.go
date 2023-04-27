package main

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUsers(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/users", nil)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := getUsers(c)

	assert.NoError(t, err)
	assert.Equal(t, 200, rec.Code)

}
