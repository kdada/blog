package controllers

import (
	"blog/backend/models"
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 用户登陆注册控制器
type AccountController struct {
	BaseController
	UserService *services.UserService
}

// Login 用户登陆
func (this *AccountController) Login(params *models.Login) web.PostResult {
	var detail, err = this.UserService.Check(params.Email, params.Password)
	if err == nil {
		//设置登陆信息
		this.Context.Session.SetBool("Login", true)
		this.Context.Session.SetValue("User", detail)
	}
	return this.returnPostResult(nil, err)
}

// Logout 用户注销
func (this *AccountController) Logout() web.PostResult {
	//检查用户是否已经登陆
	var login, ok = this.Context.Session.Bool("Login")
	if ok && login {
		this.Context.Session.Die()
		return this.Context.Json(models.NewSuccessResult(nil))
	}
	return this.Context.NotFound()
}

// Register 用户注册
func (this *AccountController) Register(params *models.Register) web.PostResult {
	return this.returnPostResult(nil, this.UserService.Add(params.Email, params.Name, params.Password))
}
