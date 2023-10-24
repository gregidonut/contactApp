package application

// implementing the appInterFace for logging and accessing some fields

func (app *Application) Debug(msg string, args ...any) {
	app.Logger.Debug(msg, args...)
}

func (app *Application) Info(msg string, args ...any) {
	app.Logger.Info(msg, args...)
}

func (app *Application) Warning(msg string, args ...any) {
	app.Logger.Warn(msg, args...)
}

func (app *Application) Error(msg string, args ...any) {
	app.Logger.Error(msg, args...)
}
