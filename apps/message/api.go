package message

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/mpush/apps/middleware"
)

type Handler struct{}

func init() {
	ioc.Container.Registry(ioc.HandlerNamespace, &Handler{})
}

func (h *Handler) Init() error {
	return nil
}

func (h *Handler) Name() string {
	return "message"
}

func (h *Handler) Registry(r gin.IRouter) {
	auth := middleware.NewAuth()
	r = r.Group("/message", auth.Token)
	{
		r.GET(".", h.Hello)
	}
}

func (h *Handler) Hello(c *gin.Context) {
	c.JSON(200, "hello")
}
