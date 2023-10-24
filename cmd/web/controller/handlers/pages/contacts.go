package pages

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Contacts(w http.ResponseWriter, r *http.Request, app *application.Application) {
	search := r.URL.Query().Get("q")
	app.Debug(fmt.Sprintf("search: %s", search))
}
