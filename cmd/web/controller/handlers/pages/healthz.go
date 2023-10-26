package pages

import (
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request, app *application.Application) {
	app.Info("app was probed for health")
	app.Info("responding with status ok")
	w.WriteHeader(http.StatusOK)
}
