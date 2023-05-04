package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	mongoClient := getDatabase("mongodb://localhost:27017", "bookstore")
	booksRepository := NewBooksRepository(mongoClient)

	handler := BookHandler(booksRepository)

	e.GET("/books", handler.getBooks)
	e.POST("/books", handler.createBook)

	e.Logger.Fatal(e.Start(":8080"))

}
