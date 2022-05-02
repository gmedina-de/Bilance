package decorators

import (
	"genuine/controllers"
)

type Decorator interface {
	Decorate(req controllers.Request, res controllers.Response)
}
