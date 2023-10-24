package application

import (
	"github.com/gregidonut/contactApp/cmd/web/model"
	"log/slog"
	"net/http"
	"os"
)

// Application is the main application object
type Application struct {
	Logger *slog.Logger
	Model  *model.Model
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

	m, err := model.NewModel(payload)
	if err != nil {
		payload.Logger.Error(err.Error())
	}

	payload.Model = m

	return payload, nil
}

func (app *Application) CatchHandlerErr(w http.ResponseWriter, err error, status int) {
	if err == nil {
		goto logToSLog
	}

	http.Error(w, err.Error(), status)

logToSLog:
	app.Logger.Error("controller error", slog.With(err))
}
