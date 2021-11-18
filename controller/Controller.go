package controller

import (
	"Bilance/service"
)

type Controller interface {
	Routing(router service.Router)
}
