package controller

import (
	"Bilance/service/router"
)

type Controller interface {
	Routing(router router.Router)
}
