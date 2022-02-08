package log

type Log interface {
	Emergency(format string, v ...interface{})
	Alert(format string, v ...interface{})
	Critical(format string, v ...interface{})
	Error(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Notice(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
}
