package service

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
