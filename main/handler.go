package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type bookHandler struct {
	BooksRepository *BooksRepository
}

func BookHandler(booksRepository *BooksRepository) *bookHandler {
	return &bookHandler{
		BooksRepository: booksRepository,
	}

}

func (h *bookHandler) getBooks(c echo.Context) error {

	books, _ := h.BooksRepository.GetBooks(c.Request().Context()) //TODO: error handler
	return c.JSON(http.StatusOK, books)
}
