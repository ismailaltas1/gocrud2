package main

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"log"
)

const booksTableName = "books"

type BooksRepository struct {
	db *mongo.Collection
}

type IBooksRepository interface {
	GetBooks(ctx context.Context) (b []Book, err error)
	InsertBooks(ctx context.Context, book Book) (err error)
}

func NewBooksRepository(db *mongo.Database) IBooksRepository {
	return &BooksRepository{
		db: db.Collection(booksTableName),
	}
}

func (br *BooksRepository) GetBooks(ctx context.Context) (b []Book, err error) {

	cur, err := br.db.Find(ctx, bson.M{})

	if err != nil {
		log.Fatal("Error get books connection")
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var books Book
		if err = cur.Decode(&books); err != nil {
			log.Fatal("Error get books connection")
		}
		b = append(b, books)
	}
	return b, err

}

func (br *BooksRepository) InsertBooks(ctx context.Context, book Book) (err error) {

	_, err = br.db.InsertOne(ctx, book)
	if err != nil {
		log.Fatalf("error exist while insterting book..")
	}
	return err

}
