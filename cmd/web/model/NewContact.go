package model

import (
	"context"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func (m *Model) NewContact(firstName, lastName, phoneNumber, email string) error {
	payload := new(contact.Contact)
	payload.FirstName = firstName
	payload.LastName = lastName
	payload.PhoneNumber = phoneNumber
	payload.EmailAddress = email

	if err := m.genMongoCollection(); err != nil {
		return err
	}

	insertResult, err := m.mongoCollection.InsertOne(context.TODO(), payload)
	if err != nil {
		log.Fatal(err)
	}

	payload.ID = insertResult.InsertedID.(primitive.ObjectID)
	m.Contacts[payload.ID] = payload

	m.app.Info("finished creating Contact instance", "contact fields", payload)
	return nil
}
