package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"html/template"
	"net/http"
)

func ContactsEdit(w http.ResponseWriter, r *http.Request, app *application.Application) {

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/edit.gohtml",
	}

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
}
