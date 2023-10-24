package model

import (
	"github.com/gregidonut/contactApp/cmd/web/utils/appInterface"
)

// Model is responsible for wrapping all the model objects so that they
// can be neatly bridged over to the main application object
type Model struct {
	app appInterface.AppInterface
}

func NewModel(app appInterface.AppInterface) (*Model, error) {
	app.Debug("creating application model..")
	defer app.Debug("finished creating application model!")

	payload := new(Model)
	payload.app = app

	return payload, nil
}
