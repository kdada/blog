package backend

import (
	"blog/backend/models"
	"blog/backend/services"

	"github.com/kdada/tinygo/meta"
	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

// 管理员登陆过滤器
type AdminFilter struct {
}

func (this *AdminFilter) Filter(context router.RouterContext) bool {
	var c, ok = context.(*web.Context)
	if ok {
		var login, ok = c.Session.Bool("Login")
		if ok && login {
			obj, ok := c.Session.Value("User")
			if ok {
				user, ok := obj.(*models.UserDetail)
				if ok {
					service, ok := meta.GlobalValueContainer.Value((*services.UserService)(nil))
					if ok {
						userService, ok := service.(*services.UserService)
						if ok {
							var admin, err = userService.IsAdmin(user.Id)
							return err == nil && admin
						}
					}
				}
			}
		}

	}
	return false
}

// 登陆过滤器
type LoginFilter struct {
}

func (this *LoginFilter) Filter(context router.RouterContext) bool {
	var c, ok = context.(*web.Context)
	if ok {
		var login, ok = c.Session.Bool("Login")
		return ok && login
	}
	return false
}
