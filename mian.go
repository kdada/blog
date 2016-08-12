package main

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/sql"
	"github.com/kdada/tinygo/web"
	_ "github.com/lib/pq"
)

// prepare 在启动前做一些准备操作
func prepare() {
	// 注册postgresql数据库 127.0.0.1:5432
	var e = sql.RegisterDefaultDB("postgres", "user=blog password=blog dbname=blog host=127.0.0.1 sslmode=disable", 100)
	if e != nil {
		panic(e)
	}

	// 添加apk文件的MIME类型
	mime.AddExtensionType(".apk", "application/vnd.android.package-archive")
}

func main() {
	//准备处理
	prepare()

	//创建web App
	var app, err = web.NewWebApp(filepath.Dir(os.Args[0]), "web.cfg", Router())
	if err != nil {
		panic(err)
	}
	//事件处理
	app.Processor.Event = new(Event)
	tinygo.AddApp(app)
	tinygo.Run()
}
