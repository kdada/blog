package main

import (
	"blog/controller"
	"blog/model"
	"mime"
	"os"
	"path/filepath"

	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

func main() {
	mime.AddExtensionType(".apk", "application/vnd.android.package-archive")
	var app, err = web.NewWebApp(filepath.Dir(os.Args[0]), "web.cfg", Router())
	if err != nil {
		panic(err)
	}
	app.Processor.Event = new(Event)
	tinygo.AddApp(app)
	tinygo.Run()
}

func Router() router.Router {
	var root = web.NewRootRouter()
	root.AddChild(web.NewControllerRouter(new(controller.HomeController)))
	root.AddChild(web.NewControllerRouter(new(controller.UserController)))
	return root
}

type Event struct {
	web.DefaultHttpProcessorEvent
}

// 每次出现一个新请求的时候触发
func (this *Event) Request(processor *web.HttpProcessor, context *web.Context) {
	processor.Logger.Debug(context.Segments())
}

// 出现错误时触发
func (this *Event) Error(processor *web.HttpProcessor, context *web.Context, err error) {
	if context != nil {
		var err = context.WriteResult(context.Api(model.NewFailureResult(1, err.Error())))
		if err != nil {
			processor.Logger.Error(err)
		}
	}
	processor.Logger.Error(err)
}
