package main

import (
	"fmt"
	"github.com/gregidonut/contactApp/cmd/web/controller/application"
	"github.com/gregidonut/contactApp/cmd/web/controller/handlers/routes"
	"log"
	"net/http"
)

const (
	DEFAULT_PORT = ":7777"
)

func main() {
	mainAppObj, err := application.NewApplication()
	if err != nil {
		log.Fatal(err)
	}
	mux := routes.Routes(mainAppObj)

	srv := &http.Server{
		Addr:    DEFAULT_PORT,
		Handler: mux,
		//ErrorLog: mainAppObj.Logger,
	}

	mainAppObj.Logger.Info(fmt.Sprintf("Starting server on %s", DEFAULT_PORT))
	err = srv.ListenAndServe()
	mainAppObj.Logger.Error(err.Error())
}
