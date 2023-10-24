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
	mainAppObj, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}

	mainAppObj.Logger.Info("starting FileServer at /static")
	fileServer := http.FileServer(http.Dir(paths.STATIC))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	for endpoint, fn := range map[string]func(w http.ResponseWriter, r *http.Request, app *application.Application){
		"/":         pages.Index,
		"/contacts": pages.Contacts,
	} {

		// these next four lines are the result of implementing a monkeypatch to any
		// HandleFunc we will create(or declared in the above for loop) since we
		// now rely on the monkey patch to expose more of the mainAppObj behavior to slog
		handler := mainAppObj.NewHandlerFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
			fn(w, r, mainAppObj)
		})
		mux.HandleFunc(fmt.Sprintf("%s", endpoint), handler.HandlerFunc())
	}

	mainAppObj.Logger.Debug(fmt.Sprintf("Starting server on %s", DEFAULT_PORT))
	err = http.ListenAndServe(DEFAULT_PORT, mux)
	mainAppObj.Logger.Error(err.Error())
}
