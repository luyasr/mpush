package user

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
	return Name
}

func (h *Handler) Registry(r gin.IRouter) {
	r = r.Group("/user")
	{
		r.POST(".", h.Create)
		r.GET(".", h.Find)
	}
}

// Create 创建用户
// @Summary 创建用户
// @Description 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param Object body CreateReq false "创建用户请求参数"
// @Success 200 {object} User
// @Router /api/v1/user [post]
func (h *Handler) Create(c *gin.Context) {
	req := new(CreateReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	user, err := h.controller.Create(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, user)
}

// Find 查询用户
// @Summary 查询用户
// @Description 查询用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param Object body FindReq false "查询用户请求参数"
// @Success 200 {object} User
// @Router /api/v1/user [get]
func (h *Handler) Find(c *gin.Context) {
	req := new(FindReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	user, err := h.controller.Find(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, user)
}
