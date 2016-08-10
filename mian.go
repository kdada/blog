package main

import (
	"blog/controllers"
	"blog/models"
	"mime"
	"os"
	"path/filepath"

	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/sql"
	"github.com/kdada/tinygo/web"
	_ "github.com/lib/pq"
)

func main() {
	var e = sql.RegisterDefaultDB("postgres", "user=blog dbname=blog host=127.0.0.1 sslmode=disable", 100)
	if e != nil {
		panic(e)
	}
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
	root.AddChild(web.NewControllerRouter(new(controllers.HomeController)))
	root.AddChild(web.NewControllerRouter(new(controllers.UserController)))
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
		var err = context.WriteResult(context.Api(models.NewFailureResult(1, err.Error())))
		if err != nil {
			processor.Logger.Error(err)
		}
	}
	processor.Logger.Error(err)
}
