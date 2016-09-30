package controllers

import (
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 用户管理控制器
type UserController struct {
	BaseController
	UserService *services.UserService
}

// List 显示用户列表,每页10行
func (this *UserController) List(params struct {
	Page int `!;>0` //页码
}) web.PostResult {
	return this.returnPostResult(this.UserService.ListAll(params.Page, pageCount))
}

// ListNum 返回页数
func (this *UserController) ListNum() web.PostResult {
	var c, err = this.UserService.UserNum()
	var result interface{} = nil
	if err == nil {
		result = map[string]int{
			"Count": calculatePages(c),
		}
	}
	return this.returnPostResult(result, err)
}

// Ban 禁止登陆
func (this *UserController) Ban(params struct {
	User int `!;>0` //用户id
}) web.PostResult {
	return this.returnPostResult(nil, this.UserService.Ban(params.User, "管理员操作"))
}

// Unban 允许登陆
func (this *UserController) Unban(params struct {
	User int `!;>0` //用户id
}) web.PostResult {
	return this.returnPostResult(nil, this.UserService.Unban(params.User))
}
