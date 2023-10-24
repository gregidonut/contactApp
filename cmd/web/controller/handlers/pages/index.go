package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, app *application.Application) {
	//w.Write([]byte(`<h1>hello world!</h1>`))

	http.Redirect(w, r, "/contacts", http.StatusPermanentRedirect)
}
