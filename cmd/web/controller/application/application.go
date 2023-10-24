package application

import (
	"github.com/gregidonut/contactApp/cmd/web/model"
	"log/slog"
	"net/http"
	"os"
)

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

// Application is the main application object
type Application struct {
	Logger *slog.Logger
	model  *model.Model
}

func (app *Application) NewHandlerFunc(
	name string,
	handlerFuncRef func(w http.ResponseWriter, r *http.Request),
) *HandlerFuncWrapper {
	return &HandlerFuncWrapper{
		app:            app,
		name:           name,
		handlerFuncRef: handlerFuncRef,
	}
}

func NewApplication() (*Application, error) {
	payload := new(Application)
	options := slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	handler := slog.NewJSONHandler(os.Stdout, &options)
	payload.Logger = slog.New(handler)

	return payload, nil
}

func (app *Application) catchHandlerErr(w http.ResponseWriter, err error, status int) {
	if err == nil {
		goto logToSLog
	}

	http.Error(w, err.Error(), status)

logToSLog:
	app.Logger.Error("controller error", slog.With(err))
}

// implementing the appInterFace for logging and accessing some fields

func (app *Application) Debug(msg string) {
	app.Logger.Debug(msg)
}

func (app *Application) Info(msg string) {
	app.Logger.Info(msg)
}

func (app *Application) Warning(msg string) {
	app.Logger.Warn(msg)
}

func (app *Application) Error(msg string) {
	app.Logger.Error(msg)
}
