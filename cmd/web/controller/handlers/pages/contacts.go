package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/model"
	"html/template"
	"net/http"
	"strings"
)

func Contacts(w http.ResponseWriter, r *http.Request, app *application.Application) {
	search := r.URL.Query()["q"]
	app.Logger.Info("logging query strings...", "query strings for 'q'", search)
	searchMatches, err := app.Model.SearchContacts(search...)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	app.Logger.Info("logging search matches as set...", "searchMatches", searchMatches)

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/index.gohtml",
		"./ui/html/components/searchForm.gohtml",
		"./ui/html/components/contacts.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", struct {
		LastSearchQuery string
		SearchMatches   model.ContactsSet
	}{
		LastSearchQuery: strings.Join(search, ", "),
		SearchMatches:   searchMatches,
	})

	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
