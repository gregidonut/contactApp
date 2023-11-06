package pages

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/model/contact"
	"html/template"
	"net/http"
)

func ContactsNew(w http.ResponseWriter, r *http.Request, app *application.Application) {
	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/new.gohtml",
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		app.Logger.Warn(fmt.Sprintf("unhandled, Method type: %s; doing nothing...", r.Method))
		return
	}

	newContact := contact.Contact{}
	if r.Method == http.MethodPost {
		app.Logger.Info("parsing form since received POST method..")

		if err := r.ParseForm(); err != nil {
			app.CatchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}

		app.Logger.Info("creating new Contact instance...")
		newCont, err := app.Model.NewContact(
			r.Form.Get("first-name"),
			r.Form.Get("last-name"),
			r.Form.Get("phone"),
			r.Form.Get("email"),
		)
		if err != nil {
			app.CatchHandlerErr(w, err, http.StatusInternalServerError)
			return
		}

		newContact = *newCont
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", newContact)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
