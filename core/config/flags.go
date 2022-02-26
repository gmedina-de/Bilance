package config

import (
	"flag"
)

var serverPort = flag.Int("port", 8080, "application port")
var logLevel = flag.Int("log", 4, "log level where 0 is critical and 4 is debug")
var viewDirectory = flag.String("views", "views", "directory where views are stored")

func ServerPort() int {
	return *serverPort
}

func LogLevel() int {
	return *logLevel
}

func ViewDirectory() string {
	return *viewDirectory
}
