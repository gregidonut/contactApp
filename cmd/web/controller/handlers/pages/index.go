package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, app *application.Application) {
	//if r.URL.Path != "/" {
	//	http.NotFound(w, r)
	//	return
	//}

	//files := []string{
	//	"./ui/html/base.gohtml",
	//	"./ui/html/pages/index.gohtml",
	//}
	//
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.CatchHandlerErr(w, err, http.StatusInternalServerError)
	//	return
	//}
	//
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.CatchHandlerErr(w, err, http.StatusInternalServerError)
	//	return
	//}
	//

	w.Write([]byte(`<h1>hello world!</h1>`))
}
