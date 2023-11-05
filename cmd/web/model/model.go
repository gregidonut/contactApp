package model

import (
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"github.com/gregidonut/contactApp/cmd/web/utils/appInterface"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContactsSet map[primitive.ObjectID]*contact.Contact

// Model is responsible for wrapping all the model objects so that they
// can be neatly bridged over to the main application object
type Model struct {
	app             appInterface.AppInterface
	Contacts        ContactsSet
	mongoCollection *mongo.Collection
}

func NewModel(app appInterface.AppInterface) (*Model, error) {
	app.Info("creating application model..")
	defer app.Info("finished creating application model!")

	payload := new(Model)
	payload.app = app

	payload.Contacts = ContactsSet{}

	if err := payload.getContacts(); err != nil {
		return payload, err
	}

	return payload, nil
}
