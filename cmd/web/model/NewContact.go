package model

import (
	"context"
	"errors"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m *Model) NewContact(firstName, lastName, phoneNumber, email string) (*contact.Contact, error) {
	payload := new(contact.Contact)

	payload.FirstName = firstName
	payload.LastName = lastName
	payload.PhoneNumber = phoneNumber
	payload.EmailAddress = email

	for _, f := range []string{
		firstName,
		lastName,
		phoneNumber,
		email,
	} {
		if f == "" {
			return payload, errors.New("required field is empty")
		}
	}
	if len([]byte(phoneNumber)) != 10 {
		return payload, errors.New("phone Number field length is not 10")
	}

	if err := m.genMongoCollection(); err != nil {
		return payload, err
	}

	insertResult, err := m.mongoCollection.InsertOne(context.TODO(), payload)
	if err != nil {
		return payload, err
	}

	payload.ID = insertResult.InsertedID.(primitive.ObjectID)
	m.Contacts[payload.ID] = payload

	m.app.Info("finished creating Contact instance", "contact fields", payload)
	return payload, nil
}
