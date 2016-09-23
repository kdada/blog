package controllers

import (
	"blog/backend/models"
	"blog/backend/services"

	"github.com/kdada/tinygo/web"
)

// 每页数量
const pageCount = 10

// 用户管理控制器
type UserManagerController struct {
	Context     *web.Context
	UserService *services.UserService
}

// List 显示用户列表,每页10行
func (this *UserManagerController) List(params struct {
	Page int `!;>0` //页码
}) web.PostResult {
	var r, err = this.UserService.ListAll(params.Page, pageCount)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(r))
}

// ListNum 返回页数
func (this *UserManagerController) ListNum() web.PostResult {
	var r, err = this.UserService.UserNum()
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	var pages = r / pageCount
	if r%pageCount > 0 {
		pages++
	}
	return this.Context.Json(models.NewSuccessResult(map[string]int{
		"Count": pages,
	}))
}

// Ban 禁止登陆
func (this *UserManagerController) Ban(params struct {
	User int `!;>0` //用户id
}) web.PostResult {
	var err = this.UserService.Ban(params.User)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}

// Unban 允许登陆
func (this *UserManagerController) Unban(params struct {
	User int `!;>0` //用户id
}) web.PostResult {
	var err = this.UserService.Unban(params.User)
	if err != nil {
		this.Context.Processor.Logger.Error(err)
		return this.Context.Json(models.NewErrorResult(err))
	}
	return this.Context.Json(models.NewSuccessResult(nil))
}
