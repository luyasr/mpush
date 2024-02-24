package cmd

import (
	"github.com/luyasr/gaia/app"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/log"
	"github.com/luyasr/mpush/apps/server"
)

func run() {
	application := app.New(app.Server(server.NewHttpServer()))
	application.Run()
}

func Init() {
	if err := ioc.Container.Init(); err != nil {
		log.Fatal(err.Error())
	}
}
