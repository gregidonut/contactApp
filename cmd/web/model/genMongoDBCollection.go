package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const MONGO_URI = "MONGO_URI"

func (m *Model) genMongoCollection() error {
	mongoDBColVals := struct {
		uri        string
		database   string
		collection string
	}{
		uri:        os.Getenv(MONGO_URI),
		database:   "contactapp",
		collection: "contacts",
	}

	clientOptions := options.Client().ApplyURI(mongoDBColVals.uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		return err
	}

	m.mongoCollection = client.Database(mongoDBColVals.database).Collection(mongoDBColVals.collection)

	return nil
}
