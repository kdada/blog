package backend

import (
	"blog/backend/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// routers 创建应用路由
func routers() router.Router {
	var root = web.NewFuncRouter("", controllers.Index)
	var manager = web.NewSpaceRouter("manager")
	manager.AddChild(web.NewMutableFuncRouter("", controllers.Manager))
	root.AddChild(manager)
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryController)))
	root.AddChild(web.NewControllerRouter(new(controllers.FileController)))
	root.AddChild(web.NewControllerRouter(new(controllers.ArticleController)))
	return root
}
