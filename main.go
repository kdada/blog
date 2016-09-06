package main

import (
	"blog/backend"

	"github.com/kdada/tinygo"
)

// main 启动App
func main() {
	tinygo.AddApp(backend.NewApp())

	tinygo.Run()
}
