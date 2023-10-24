package model

import (
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"github.com/gregidonut/contactApp/cmd/web/utils/appInterface"
)

// Model is responsible for wrapping all the model objects so that they
// can be neatly bridged over to the main application object
type Model struct {
	app      appInterface.AppInterface
	contacts []*contact.Contact
}

func NewModel(app appInterface.AppInterface) (*Model, error) {
	app.Debug("creating application model..")
	defer app.Debug("finished creating application model!")

	payload := new(Model)
	payload.app = app

	return payload, nil
}

func (m *Model) SearchContacts(filters ...string) ([]*contact.Contact, error) {
	var payload []*contact.Contact
	if (len(filters) == 1 && filters[0] == "") || len(filters) == 0 || filters == nil {
		m.app.Warning("search contacts called without args returning all contacts")
		return m.contacts, nil
	}

	//for _, filter := range filters {
	//	for _, cont := range m.contacts {
	//
	//	}
	//}

	return payload, nil
}
