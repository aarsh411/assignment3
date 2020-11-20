package main

import (
	"github.com/aarsh411/Todo/app"
	"github.com/aarsh411/Todo/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}