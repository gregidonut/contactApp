package pages

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Contacts(w http.ResponseWriter, r *http.Request, app *application.Application) {
	search := r.URL.Query().Get("q")
	searchMatches, err := app.Model.SearchContacts(search)
	if err != nil {
		app.CatchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	app.Debug(fmt.Sprintf("printing out a list of matches"))
	for _, match := range searchMatches {
		app.Debug(fmt.Sprintf("%#v", match))
	}

}
