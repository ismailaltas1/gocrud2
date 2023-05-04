package main

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/context"
	"net/http"
)

type bookHandler struct {
	BooksRepository IBooksRepository
}

func BookHandler(booksRepository IBooksRepository) *bookHandler {
	return &bookHandler{
		BooksRepository: booksRepository,
	}

}

func (h *bookHandler) getBooks(c echo.Context) error {

	books, _ := h.BooksRepository.GetBooks(context.Background()) //TODO: error handler
	return c.JSON(http.StatusOK, books)
}

func (h *bookHandler) createBook(c echo.Context) error {
	var book Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	err := h.BooksRepository.InsertBooks(context.Background(), book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, book)

}
