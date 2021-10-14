package main

import (
	"github.com/flatcar-linux/flog/app"
	"github.com/flatcar-linux/flog/config"
)

func main() {
	config := config.GetConfig()
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
