package main

import (
	"blog/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// 创建路由
func Router() router.Router {
	var root = web.NewRootRouter()

	root.AddChild(web.NewFuncRouter(`index.html`, controllers.Index))

	root.AddChild(web.NewFuncRouter(`a{Id=\d+}.html`, controllers.Article))

	var category = web.NewFuncRouter(`c{Id=\d+}`, controllers.Category)
	category.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Category))
	root.AddChild(category)

	root.AddChild(web.NewStaticRouter("img", "app/img"))

	var manager = web.NewSpaceRouter("manager")
	manager.AddChild(web.NewMutableFuncRouter("index", controllers.Manager))
	root.AddChild(manager)

	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	return root
}
