package db

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client            *mongo.Client
	ClientInstanceErr error
	mongoOnce         sync.Once
)

const (
	CONNECTIONSTRIING = "mongodb+srv://suriya:ayirus@cluster0.kicok.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	DB                = "bookstore"
	ISSUES            = ""
)

func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRIING)

		clientInstance, err := mongo.Connect(context.TODO(), clientOptions)

		if err != nil {
			ClientInstanceErr = err
		}

		err = clientInstance.Ping(context.TODO(), nil)

		if err != nil {
			ClientInstanceErr = err
		}
		Client = clientInstance
	})

	return Client, ClientInstanceErr
}
