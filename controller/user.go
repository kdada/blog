package controller

import (
	"blog/model"
	"blog/service"

	"github.com/kdada/tinygo/web"
)

type UserController struct {
	Context *web.Context
}

func (this *UserController) Login(m *model.UserLogin) web.PostResult {
	var s = new(service.UserService)
	var userInfo, err = s.Login(m)
	if err != nil {
		return this.Context.Api(model.NewFailureResult(1, err.Error()))
	}
	this.Context.Session.SetBool("Login", true)
	this.Context.Session.SetValue("User", userInfo)
	return this.Context.Api(model.NewSuccessResult(userInfo))
}

func (this *UserController) Register(m *model.UserRegister) web.PostResult {
	var s = new(service.UserService)
	var err = s.Register(m)
	if err != nil {
		return this.Context.Api(model.NewFailureResult(1, err.Error()))
	}
	return this.Context.Api(model.NewSuccessResult(nil))
}
