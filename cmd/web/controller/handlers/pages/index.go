package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, app *application.Application) {
	app.Debug(`redirecting to "/contacts" endpoint`)
	http.Redirect(w, r, "/contacts", http.StatusPermanentRedirect)
}
