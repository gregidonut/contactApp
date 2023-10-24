package model

import (
	"encoding/json"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"github.com/gregidonut/contactApp/cmd/web/utils/appInterface"
	"os"
	"reflect"
	"strings"
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

func (m *Model) SearchContacts(filters ...string) (map[int]*contact.Contact, error) {
	m.app.Info("running SearchContacts method...")
	defer m.app.Info("finished running SearchContacts method!")

	//{{ mocking the generation of contacts;
	//   in the real world, this search would probably be done by searching from the result
	//   of a database query or better yet, maybe the orm has a search api that can be called
	jsonData, err := os.ReadFile("testingAssets/contactData.json")
	if err != nil {
		return nil, err
	}

	var aux []*contact.Contact
	err = json.Unmarshal(jsonData, &aux)
	if err != nil {
		return nil, err
	}

	for _, i := range aux {
		m.contacts[i.ID] = i
	}
	//}}

	payload := map[int]*contact.Contact{}

	if (len(filters) == 1 && filters[0] == "") || len(filters) == 0 || filters == nil {
		m.app.Warning("search contacts called without args returning all contacts")
		return m.contacts, nil
	}

	for _, filter := range filters {
		for _, cont := range m.contacts {
			v := reflect.ValueOf(*cont)
			for i := 0; i < v.NumField(); i++ {
				field := v.Field(i)
				if field.Kind() != reflect.String {
					continue
				}
				fieldValue := field.String()
				if !strings.Contains(strings.ToLower(fieldValue), strings.ToLower(filter)) {
					continue
				}
				payload[cont.ID] = cont
			}
		}
	}

	return payload, nil
}
