package channel

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/transport/http/response"
	"strconv"
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
	return Name
}

func (h *Handler) Registry(r gin.IRouter) {
	group := r.Group("/channel")
	{
		group.POST(".", h.Create)
		group.PUT(":id", h.Update)
		group.DELETE(":id", h.Delete)
		group.POST(":id", h.Query)
	}
}

// Create 创建频道
// @Summary 创建频道
// @Description 创建频道
// @Tags channel
// @Accept json
// @Produce json
// @Param Object body CreateReq true "创建频道请求参数"
// @Success 200 {object} Channel
// @Router /api/v1/channel [post]
func (h *Handler) Create(c *gin.Context) {
	req := new(CreateReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	channel, err := h.controller.Create(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, channel)
}

// Update 更新频道
// @Summary 更新频道
// @Description 更新频道
// @Tags channel
// @Accept json
// @Produce json
// @Param Object body UpdateReq true "更新频道请求参数"
// @Success 200
// @Router /api/v1/channel/{id} [PUT]
func (h *Handler) Update(c *gin.Context) {
	req := new(UpdateReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	if err := h.controller.Update(c, req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, nil)
}

// Delete 删除频道
// @Summary 删除频道
// @Description 删除频道
// @Tags channel
// @Accept json
// @Produce json
// @Param id path int true "频道ID"
// @Success 200
// @Router /api/v1/channel/delete/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)

	if err := h.controller.Delete(c, parseInt); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, nil)
}

// Query 查询频道
// @Summary 查询频道
// @Description 查询频道
// @Tags channel
// @Accept json
// @Produce json
// @Param Object body QueryReq true "查询频道请求参数"
// @Success 200 {object} Channels
// @Router /api/v1/channel/{id} [post]
func (h *Handler) Query(c *gin.Context) {
	req := new(QueryReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	channels, err := h.controller.Query(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, channels)
}
