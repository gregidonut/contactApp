package application

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
