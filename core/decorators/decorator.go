package decorators

import "genuine/core/controllers"

type Decorator interface {
	Decorate(req controllers.Request, res controllers.Response)
}
