package appInterface

type AppInterface interface {
	Debug(string)
	Info(string)
	Warning(string)
	Error(string)
}
