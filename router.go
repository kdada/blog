package main

import (
	"blog/controllers"

	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// 创建路由
func Router() router.Router {
	//根路由
	var root = web.NewRootRouter()
	//首页路由
	root.AddChild(web.NewFuncRouter(`index.html`, controllers.Index))
	root.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Index))
	//文章路由
	root.AddChild(web.NewFuncRouter(`a{Id=\d+}.html`, controllers.Article))
	//分类列表路由
	var category = web.NewFuncRouter(`c{Id=\d+}`, controllers.Category)
	category.AddChild(web.NewFuncRouter(`p{Page=\d+}.html`, controllers.Category))
	root.AddChild(category)
	//图片路由
	root.AddChild(web.NewStaticRouter("img", "app/img"))
	//管理后台路由
	var manager = web.NewSpaceRouter("manager")
	manager.AddChild(web.NewMutableFuncRouter("index", controllers.Manager))
	root.AddChild(manager)
	//用户服务路由
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
	//评论路由
	root.AddChild(web.NewControllerRouter(new(controllers.ReplyController)))
	//分类路由
	root.AddChild(web.NewControllerRouter(new(controllers.CategoryController)))
	return root
}
