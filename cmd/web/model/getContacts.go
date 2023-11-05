package model

import (
	"context"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const MONGO_URI = "MONGO_URI"

var ctx = context.TODO()

func (m *Model) getContacts() error {
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
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	cur, err := client.Database(mongoDBColVals.database).Collection(mongoDBColVals.collection).Find(ctx, bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result contact.Contact
		err := cur.Decode(&result)
		if err != nil {
			return err
		}
		m.app.Debug("logging contents of collection...", "contactapp.contacts", result)
		m.contacts[result.ID] = &result
	}
	m.app.Debug("logging contents of contacts in model", "[]contacts", m.contacts)

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}
