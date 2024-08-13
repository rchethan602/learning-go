package controller

import (
	"bookstore/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://atlas-go:zaq21wsx@cluster0.qsh5ddp.mongodb.net/?retryWrites=true&w=majority"

const dbName = "bookstore"

const colName = "books"

//MOST import as connection

var collection *mongo.Collection

//connect with mongodb

func init() {
	//clientOption := options.Client().ApplyURI(connectionString)

	//connect to database

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))

	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("bookstore").Collection("books")
}

func getAllBooks() []primitive.M {

	cur, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var books []primitive.M

	for cur.Next(context.Background()) {
		var book bson.M

		err := cur.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}
	defer cur.Close(context.Background())
	return books
}

func insertAbook(book models.Book) {
	inserted, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(inserted.InsertedID)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	allBooks := getAllBooks()
	json.NewEncoder(w).Encode(allBooks)
}

func InsertAbook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var book models.Book

	json.NewDecoder(r.Body).Decode(&book)

	insertAbook(book)

	json.NewEncoder(w).Encode(book)
}
