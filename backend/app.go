package backend

import (
	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/web"
)

// NewApp 创建App
func NewApp() tinygo.App {
	var app, err = web.NewWebApp("./", "config/web.cfg", routers())
	if err != nil {
		panic(err)
	}
	return app
}
