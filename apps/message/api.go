package message

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/transport/http/response"
)

type Handler struct {
	controller *Controller
}

func init() {
	ioc.Container.Registry(ioc.HandlerNamespace, &Handler{})
}

func (h *Handler) Init() error {
	h.controller = ioc.Container.Get(ioc.ControllerNamespace, Name).(*Controller)
	return nil
}

func (h *Handler) Name() string {
	return "message"
}

func (h *Handler) Registry(r gin.IRouter) {
	r = r.Group("/message")
	{
		r.POST(".", h.Send)
		r.GET(".", h.Read)
	}
}

func (h *Handler) Send(c *gin.Context) {
	req := new(ClientSendReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}
	if err := h.controller.ClientSend(c, req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, nil)
}

func (h *Handler) Read(c *gin.Context) {
	h.controller.Read()
}
