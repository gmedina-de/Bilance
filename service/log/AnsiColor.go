package log

type AnsiColor string

const (
	Reset    AnsiColor = "\x1b[0m"
	FgBlack  AnsiColor = "\x1b[30m"
	FgRed    AnsiColor = "\x1b[31m"
	FgGreen  AnsiColor = "\x1b[32m"
	FgYellow AnsiColor = "\x1b[33m"
	FgBlue   AnsiColor = "\x1b[34m"
	FgPurple AnsiColor = "\x1b[35m"
	FgCyan   AnsiColor = "\x1b[36m"
	FgWhite  AnsiColor = "\x1b[37m"
	BgBlack  AnsiColor = "\x1b[40m"
	BgRed    AnsiColor = "\x1b[41m"
	BgGreen  AnsiColor = "\x1b[42m"
	BgYellow AnsiColor = "\x1b[43m"
	BgBlue   AnsiColor = "\x1b[44m"
	BgPurple AnsiColor = "\x1b[45m"
	BgCyan   AnsiColor = "\x1b[46m"
	BgWhite  AnsiColor = "\x1b[47m"
)
