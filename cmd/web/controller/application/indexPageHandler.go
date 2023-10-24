package application

import (
	"html/template"
	"net/http"
)

func (app *Application) IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.gohtml",
		"./ui/html/pages/index.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.catchHandlerErr(w, err, http.StatusInternalServerError)
		return
	}
}
