package backend

import (
	"blog/backend/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// routers 创建应用路由
func routers() router.Router {
	//首页路由
	var root = web.NewFuncRouter("", controllers.Index)
	root.AddChild(web.NewFuncRouter(`p{Page=\d+?}.html`, controllers.Index))

	//分类列表路由
	var category = web.NewFuncRouter(`c{Category=\d+}`, controllers.Category)
	category.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Category))
	root.AddChild(category)

	//文章路由
	root.AddChild(web.NewFuncRouter(`a{Article=\d+?}.html`, controllers.Article))

	//管理页面
	var manager = web.NewSpaceRouter("manager")
	manager.AddChild(web.NewMutableFuncRouter("", controllers.Manager))
	root.AddChild(manager)

	//接口
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryController)))
	root.AddChild(web.NewControllerRouter(new(controllers.FileController)))
	root.AddChild(web.NewControllerRouter(new(controllers.ArticleController)))
	return root
}
