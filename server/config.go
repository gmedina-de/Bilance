package server

import (
	"flag"
)

var webdav_port = flag.Int("webdav_port", 8081, "webdav server address")
var webdav_path = flag.String("webdav_path", "./data", "webdav data path")
