package main

import (
	"blog/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

func Router() router.Router {
	var root = web.NewRootRouter()
	root.AddChild(web.NewControllerRouter(new(controllers.HomeController)))
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	return root
}
