package main

import (
	"blog/controller"
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
	return root
}

type Event struct {
	web.DefaultHttpProcessorEvent
}

// 每次出现一个新请求的时候触发
func (this *Event) Request(processor *web.HttpProcessor, context *web.Context) {
	processor.Logger.Debug(context.Segments())
}
