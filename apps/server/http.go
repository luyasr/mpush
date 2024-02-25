package server

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/middleware/log"
	"github.com/luyasr/gaia/middleware/recovery"
	"github.com/luyasr/gaia/transport/http"
	"github.com/luyasr/mpush/config"
)

func logMode(mode string) string {
	var m string
	switch mode {
	case "debug":
		m = "debug"
	case "release":
		m = "release"
	case "test":
		m = "test"
	default:
		m = "release"
	}

	return m
}

func NewHttpServer() *http.Server {
	gin.SetMode(logMode(config.Cfg.Http.Mode))
	r := gin.New()
	r.Use(log.New().GinLogger(), recovery.New().GinRecovery(true))
	apiV1 := r.Group("/api/v1")

	ioc.Container.GinIRouterRegistry(apiV1)

	server := http.NewServer(http.Address(config.Cfg.Http.Address()), http.Handler(r))

	return server
}
