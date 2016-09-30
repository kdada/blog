package backend

import (
	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/sql"
	"github.com/kdada/tinygo/web"
	_ "github.com/lib/pq"
)

// NewApp 创建App
func NewApp() tinygo.App {
	sql.RegisterDefaultDB("postgres", "user=blog password=blog dbname=blog sslmode=disable", 10)
	var app, err = web.NewWebApp("./", "config/web.cfg", routers())
	if err != nil {
		panic(err)
	}
	app.Processor.Event = new(ProcessorEvent)
	return app
}
