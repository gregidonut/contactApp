package CRUD

import (
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

	app.Logger.Info("submitting form since received POST method..")

	return
}
