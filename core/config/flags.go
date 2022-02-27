package config

import (
	"flag"
)

var logLevel = flag.Int("log", 4, "log level where 0 is critical and 4 is debug")

func LogLevel() int {
	return *logLevel
}

var databaseLocation = flag.String("db", "./database.db", "database location")

func DatabaseLocation() string {
	return *databaseLocation
}

var serverPort = flag.Int("port", 8080, "application port")

func ServerPort() int {
	return *serverPort
}

var viewDirectory = flag.String("views", "views", "directory where views are stored")

func ViewDirectory() string {
	return *viewDirectory
}

func init() {
	flag.Parse()
}
