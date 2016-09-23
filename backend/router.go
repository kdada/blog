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
	root.AddChild(web.NewControllerRouter(new(controllers.UserManagerController)))
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryManagerController)))

	return root
}
