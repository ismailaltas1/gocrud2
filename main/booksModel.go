package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID     primitive.ObjectID `json:"_id"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
}
