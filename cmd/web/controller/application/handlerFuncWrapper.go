package application

import "net/http"

// HandlerFuncWrapper is needed to ultimately append and/or prepend logic to
// the handler functions programmatically.
// Because of this, every endpoint where HandlerFunc is called, the info.logger messages
// declared in NewHandlerFunc (which should be required before registering to the mux),
// will have these log messages. or anything added to the current HandlerFunc declaration
type HandlerFuncWrapper struct {
	app            *Application
	name           string
	handlerFuncRef func(w http.ResponseWriter, r *http.Request)
}

func (hfw *HandlerFuncWrapper) HandlerFunc() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hfw.app.Logger.Info("started running", "endpoint", hfw.name)
		defer hfw.app.Logger.Info("finished running", "endpoint", hfw.name)

		hfw.handlerFuncRef(w, r)
	}
}
