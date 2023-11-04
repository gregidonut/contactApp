package model

import (
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"github.com/gregidonut/contactApp/cmd/web/utils/appInterface"
)

// Model is responsible for wrapping all the model objects so that they
// can be neatly bridged over to the main application object
type Model struct {
	app      appInterface.AppInterface
	contacts map[int]*contact.Contact //set
}

func NewModel(app appInterface.AppInterface) (*Model, error) {
	app.Info("creating application model..")
	defer app.Info("finished creating application model!")

	payload := new(Model)
	payload.app = app
	payload.contacts = map[int]*contact.Contact{}

	return payload, nil
}
