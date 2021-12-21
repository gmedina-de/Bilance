package log

type Log interface {
	Debug(message string, parameters ...interface{})
	Info(message string, parameters ...interface{})
	Warning(message string, parameters ...interface{})
	Error(message string, parameters ...interface{})
}

type LogLevel string

const (
	Debug   LogLevel = "DBUG"
	Info    LogLevel = "INFO"
	Warning LogLevel = "WARN"
	Error   LogLevel = "ERRO"
)

func (l LogLevel) toFgColor() LogColor {
	switch l {
	case Debug:
		return FgGreen
	case Info:
		return FgBlue
	case Warning:
		return FgYellow
	case Error:
		return FgRed
	default:
		return FgGreen
	}
}

func (l LogLevel) toBgColor() LogColor {
	switch l {
	case Debug:
		return BgGreen
	case Info:
		return BgBlue
	case Warning:
		return BgYellow
	case Error:
		return BgRed
	default:
		return BgGreen
	}
}

type LogColor string

const (
	Reset    LogColor = "\x1b[0m"
	FgBlack  LogColor = "\x1b[30m"
	FgRed    LogColor = "\x1b[31m"
	FgGreen  LogColor = "\x1b[32m"
	FgYellow LogColor = "\x1b[33m"
	FgBlue   LogColor = "\x1b[34m"
	FgPurple LogColor = "\x1b[35m"
	FgCyan   LogColor = "\x1b[36m"
	FgWhite  LogColor = "\x1b[37m"
	BgBlack  LogColor = "\x1b[40m"
	BgRed    LogColor = "\x1b[41m"
	BgGreen  LogColor = "\x1b[42m"
	BgYellow LogColor = "\x1b[43m"
	BgBlue   LogColor = "\x1b[44m"
	BgPurple LogColor = "\x1b[45m"
	BgCyan   LogColor = "\x1b[46m"
	BgWhite  LogColor = "\x1b[47m"
)