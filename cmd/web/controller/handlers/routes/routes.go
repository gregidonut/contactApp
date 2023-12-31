package routes

import (
	"github.com/gorilla/mux"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/controller/handlers/health"
	"github.com/gregidonut/contactApp/cmd/web/controller/handlers/pages"
	"github.com/gregidonut/contactApp/cmd/web/utils/paths"
	"net/http"
)

// handlerFuncRef is to be used as a http.HandlerFunc but with the Application pointer passed to it
// to give it the ability to expose handler behavior with logging and also model object methods for
// a more MVC approach to the web app
type handlerFuncRef func(http.ResponseWriter, *http.Request, *application.Application)

func Routes(app *application.Application) *mux.Router {
	gorillaMux := mux.NewRouter()

	app.Logger.Info("starting FileServer at /static")
	fileServer := http.FileServer(http.Dir(paths.STATIC))
	gorillaMux.PathPrefix("/static/").Handler(http.StripPrefix("/static", fileServer))

	registerHandler := func(endpoint string, hfr handlerFuncRef) {
		gorillaMux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			app.Logger.Info("running", "endpoint", endpoint)
			defer app.Logger.Info("completed", "endpoint", endpoint)

			hfr(w, r, app)
		})
	}
	var endpointRegistry = map[string]handlerFuncRef{
		"/":                                   pages.Index,
		"/contacts":                           pages.Contacts,
		"/contacts/{id:[A-Za-z0-9]{24}}":      pages.ContactsDetails,
		"/contacts/{id:[A-Za-z0-9]{24}}/edit": pages.ContactsEdit,
		"/contacts/new":                       pages.ContactsNew,

		"/healthz": health.Healthz,
	}

	// this loop is the result of implementing a monkeypatch to any HandleFunc we
	// will create(or declared in the above for endpointRegistry var) since we
	// now rely on the monkey patch to expose more of the mainAppObj behavior to slog
	for endpoint, hfr := range endpointRegistry {
		registerHandler(endpoint, hfr)
	}
	//}}

	return gorillaMux
}
