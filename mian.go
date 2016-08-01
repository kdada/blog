package main

import (
	"os"
	"path/filepath"

	"github.com/kdada/tinygo"
	"github.com/kdada/tinygo/router"
	"github.com/kdada/tinygo/web"
)

func main() {
	start()
}

func start() {
	var app, err = web.NewWebApp(filepath.Dir(os.Args[0]), "web.cfg", Router())
	if err != nil {
		panic(err)
	}
	tinygo.AddApp(app)
	tinygo.Run()
}

func Router() router.Router {
	var root = web.NewRootRouter()
	return root
}
