package main

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/controller/handlers/pages"
	"github.com/gregidonut/contactApp/cmd/web/utils/paths"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":8080"
)

func main() {
	mux := http.NewServeMux()
	app, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	app.Logger.Info("starting FileServer at /static")
	fileServer := http.FileServer(http.Dir(paths.STATIC))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	for endpoint, fn := range map[string]func(w http.ResponseWriter, r *http.Request, app *application.Application){
		"/": pages.Index,
	} {

		// these next four lines are the result of implementing a monkeypatch to any
		// HandleFunc we will create(or declared in the above for loop) since we
		// now rely on the monkey patch to expose more of the app behavior to slog
		handler := app.NewHandlerFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, app)
		})
		mux.HandleFunc(fmt.Sprintf("%s", endpoint), handler.HandlerFunc())
	}

	app.Logger.Debug(fmt.Sprintf("Starting server on %s", DEFAULT_PORT))

	err = http.ListenAndServe(DEFAULT_PORT, mux)
	app.Logger.Error(err.Error())
}
