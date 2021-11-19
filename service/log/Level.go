package log

type Level string

const (
	Debug   Level = "DBUG"
	Info    Level = "INFO"
	Warning Level = "WARN"
	Error   Level = "ERRO"
)

func (l Level) toFgColor() AnsiColor {
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

func (l Level) toBgColor() AnsiColor {
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
