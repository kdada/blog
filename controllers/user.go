package controllers

import (
	"blog/models"
	"blog/services"

	"github.com/kdada/tinygo/web"
)

type UserController struct {
	Context *web.Context
}

func (this *UserController) Login(m *models.UserLogin, userService *services.UserService) web.PostResult {
	var userInfo, err = userService.Login(m)
	if err != nil {
		return this.Context.Api(models.NewFailureResult(1, err.Error()))
	}
	this.Context.Session.SetBool("Login", true)
	this.Context.Session.SetValue("User", userInfo)
	return this.Context.Api(models.NewSuccessResult(userInfo))
}

func (this *UserController) Register(m *models.UserRegister, userService *services.UserService) web.PostResult {
	var err = userService.Register(m)
	if err != nil {
		return this.Context.Api(models.NewFailureResult(1, err.Error()))
	}
	return this.Context.Api(models.NewSuccessResult(nil))
}
