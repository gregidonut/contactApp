package pages

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"net/http"
)

func ContactsDetails(w http.ResponseWriter, r *http.Request, app *application.Application) {
	contactId := mux.Vars(r)["id"]
	app.Logger.Debug("logging contactId from url var", "mux.Vars(r)[\"id\"]", contactId)

	objectID, err := primitive.ObjectIDFromHex(contactId)
	if err != nil {
		app.CatchHandlerErr(
			w,
			errors.New(fmt.Sprintf("unable to cast hex contactID: %s into mongodb objectid", contactId)),
			http.StatusInternalServerError,
		)
		return
	}
	app.Logger.Debug("validating existence from Model.ContactSet")
	cont, ok := app.Model.Contacts[objectID]
	if !ok {
		app.CatchHandlerErr(w, errors.New(fmt.Sprintf("contactID: %s; doesn't exist", contactId)), http.StatusBadRequest)
		return
	}
	app.Logger.Debug("validation successful!!")

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/details.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", cont)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
