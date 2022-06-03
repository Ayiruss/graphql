package users

import (
	"context"
	"net/http"
	"time"

	"github.com/Ayiruss/bookstore/graph/internal/db"
	"github.com/Ayiruss/bookstore/graph/utils/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	COLLECTION = "user"
)

var (
	collection *mongo.Collection
)

func init() {
	_, err := db.GetMongoClient()
	if err != nil {
		panic(err)
	}
	collection = db.Client.Database(db.DB).Collection(COLLECTION)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Balance   int64              `bson:"balance"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	UpdatedBy string             `bson:"updatedBy"`
}

func (user *User) CheckPasswordHash(hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hash))
}

func (user *User) GetByUserName() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"username": user.Username}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error reading User",
		}
	}

	return nil
}

// The method is not exposed as of now but can be used to increase the balance in the user's account
func (user *User) UpdateBalance(balance int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection = db.Client.Database(db.DB).Collection(COLLECTION)
	filter := bson.M{"_id": user.ID}
	update := bson.D{
		{"$set", bson.D{{"balance", balance}}},
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error updating User",
		}
	}
	upsertedID := result.UpsertedID
	print(upsertedID)
	return nil
}
