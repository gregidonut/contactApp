package model

import (
	"context"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.TODO()

func (m *Model) getContacts() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	cur, err := client.Database("test").Collection("contacts").Find(ctx, bson.M{})
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
		m.app.Debug("logging contents of collection...", "test.contacts", result)
		m.contacts[result.ID] = &result
	}
	m.app.Debug("logging contents of contacts in model", "[]contacts", m.contacts)

	if err := cur.Err(); err != nil {
		return err
	}

	return nil
}
