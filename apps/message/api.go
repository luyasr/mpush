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
		r.POST(".", h.Producer)
	}
}

// Producer 客户端发送消息
// @Summary 客户端发送消息
// @Description 客户端发送消息
// @Tags message
// @Accept json
// @Produce json
// @Param req body ProducerReq true "请求"
// @Success 200 {object} interface{} "成功"
// @Router /api/v1/message [post]
func (h *Handler) Producer(c *gin.Context) {
	req := new(ProducerReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}
	if err := h.controller.Producer(c, req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, nil)
}
