package model

import (
	"context"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *Model) getContacts() error {
	if err := m.genMongoCollection(); err != nil {
		return err
	}
	cur, err := m.mongoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var result contact.Contact
		if err = cur.Decode(&result); err != nil {
			return err
		}
		m.app.Debug("logging contents of collection...", "contactapp.contacts", result)
		m.Contacts[result.ID] = &result
	}
	m.app.Debug("logging contents of contacts in model", "[]contacts", m.Contacts)

	if err = cur.Err(); err != nil {
		return err
	}

	return nil
}
