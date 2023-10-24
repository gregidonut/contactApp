package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"html/template"
	"net/http"
)

func Contacts(w http.ResponseWriter, r *http.Request, app *application.Application) {
	search := r.URL.Query().Get("q")
	searchMatches, err := app.Model.SearchContacts(search)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	//app.Debug("printing out a list of matches")
	app.Logger.Debug("search matches as set", "searchMatches", searchMatches)

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/index.gohtml",
		"./ui/html/components/contacts.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", searchMatches)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
