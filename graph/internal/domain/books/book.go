package books

import (
	"context"
	"net/http"
	"time"

	"github.com/Ayiruss/bookstore/graph/internal/db"
	"github.com/Ayiruss/bookstore/graph/internal/helper"
	"github.com/Ayiruss/bookstore/graph/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	SellerID    string             `bson:"sellerId"`
	SellerName  string             `bson:"sellerName"`
	Price       int64              `bson:"price"`
	Status      int                `bson:"status"`
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
	UpdatedBy   string             `bson:"updatedBy"`
}

const (
	COLLECTION = "book"
	DB         = "bookstore"
)

var (
	collection *mongo.Collection
)

func (book *Book) Get() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection = db.Client.Database(DB).Collection(COLLECTION)
	filter := bson.M{"_id": book.ID}
	result := collection.FindOne(ctx, filter)
	result.Decode(book)
	return nil
}

func List() ([]*Book, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection = db.Client.Database(DB).Collection(COLLECTION)
	filter := bson.M{"status": helper.Published.EnumIndex()}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error reading book from the Table",
		}
	}
	var books_do []*Book
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var book Book
		err := cur.Decode(&book)
		if err != nil {
			return nil, &errors.MyError{
				Inner:   err,
				Message: "Error reading book from the Table",
			}
		}

		books_do = append(books_do, &book)
	}
	return books_do, nil
}

func (book *Book) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection = db.Client.Database(DB).Collection(COLLECTION)
	result, err := collection.InsertOne(ctx, book)
	if err != nil {
		return &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error writing data to the table",
		}
	}
	insertedID := result.InsertedID
	print(insertedID)
	return nil
}

func (book *Book) UpdateStatus() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection = db.Client.Database(DB).Collection(COLLECTION)
	filter := bson.M{"_id": book.ID}
	update := bson.D{
		{"$set", bson.D{{"status", book.Status}}},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error updating book in the table",
		}
	}
	upsertedID := result.UpsertedID
	print(upsertedID)
	return nil
}
