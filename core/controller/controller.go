package controller

import "homecloud/core/server"

type Controller interface {
	Routing(server server.Server)
}
