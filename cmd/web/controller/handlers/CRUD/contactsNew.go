package CRUD

import (
	"errors"
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"html/template"
	"net/http"
)

func ContactsNew(w http.ResponseWriter, r *http.Request, app *application.Application) {
	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/new.gohtml",
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		app.Logger.Warn(fmt.Sprintf("unhandled Method type: %s; doing nothing...", r.Method))
		return
	}

	if r.Method == http.MethodGet {
		app.Logger.Info("rendering form since received GET method..")

		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.CatchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.CatchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}

		return
	}

	app.Logger.Info("parsing form since received POST method..")

	if err := r.ParseForm(); err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	email := r.Form.Get("email")
	firstName := r.Form.Get("first-name")
	lastName := r.Form.Get("last-name")
	phone := r.Form.Get("phone")

	formFields := []string{email, firstName, lastName, phone}
	app.Logger.Debug("logging form fields: ", "fields", formFields)

	for _, field := range formFields {
		if field == "" {
			app.CatchHandlerErr(w, errors.New("required field empty"), http.StatusInternalServerError)
			return
		}
	}

	app.Logger.Info("creating new Contact instance...")
	if err := app.Model.NewContact(firstName, lastName, phone, email); err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
	}

	return
}
