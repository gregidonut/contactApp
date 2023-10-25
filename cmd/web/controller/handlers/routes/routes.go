package routes

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/controller/handlers/pages"
	"github.com/gregidonut/contactApp/cmd/web/utils/paths"
	"net/http"
)

func Routes(app *application.Application) *http.ServeMux {
	mux := http.NewServeMux()

	app.Logger.Info("starting FileServer at /static")
	fileServer := http.FileServer(http.Dir(paths.STATIC))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//{{
	registerHandler := func(endpoint string, handler func(http.ResponseWriter, *http.Request, *application.Application)) {
		mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			app.Logger.Info(fmt.Sprintf("registering '%s' endpoint to mux...", endpoint))
			defer app.Logger.Info(fmt.Sprintf("finished registering '%s' endpoint to mux!", endpoint))

			handler(w, r, app)
		})
	}
	endpointRegistry := map[string]func(
		w http.ResponseWriter,
		r *http.Request,
		pageRegistryAppObj *application.Application,
	){
		"/":         pages.Index,
		"/contacts": pages.Contacts,
	}

	// this loop is the result of implementing a monkeypatch to any HandleFunc we
	// will create(or declared in the above for endpointRegistry var) since we
	// now rely on the monkey patch to expose more of the mainAppObj behavior to slog
	for endpoint, handlerFuncRef := range endpointRegistry {
		//endpoint := endpnt
		registerHandler(endpoint, handlerFuncRef)
	}
	//}}

	return mux
}
