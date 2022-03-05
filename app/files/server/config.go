package server

import (
	"flag"
	"genuine/core/config"
)

var webdav_port = flag.Int("webdav_port", config.ServerPort()+1, "webdav server address")
var webdav_path = flag.String("webdav_path", "./data", "webdav data path")
