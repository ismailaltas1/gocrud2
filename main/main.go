package main

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	mongoClient := getDatabase("mongodb://localhost:27017", "bookstore")
	booksRepository := NewBooksRepository(mongoClient)

	handler := BookHandler(booksRepository)
	e := echo.New()

	e.GET("/books", handler.getBooks)

	e.Start(":8080")

}
