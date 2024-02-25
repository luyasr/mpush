package token

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
	r = r.Group("/token")
	{
		r.POST("/login", h.Login)
		r.POST("/logout", h.Logout)
		r.POST("/refresh", h.Refresh)
	}
}

// Login 登录
// @Summary 登录
// @Description 登录
// @Tags token
// @Accept json
// @Produce json
// @Param Object body LoginReq true "登录请求参数"
// @Success 200 {object} Tk
// @Router /api/v1/token/login [post]
func (h *Handler) Login(c *gin.Context) {
	req := new(LoginReq)
	if err := c.BindJSON(req); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	tk, err := h.controller.Login(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, tk)
}

// Logout 登出
// @Summary 登出
// @Description 登出
// @Tags token
// @Accept json
// @Produce json
// @Param Object body Tk true "登出请求参数"
// @Success 200
// @Router /api/v1/token/logout [post]
func (h *Handler) Logout(c *gin.Context) {
}

// Refresh 刷新
// @Summary 刷新
// @Description 刷新
// @Tags token
// @Accept json
// @Produce json
// @Param Object body Tk true "刷新请求参数"
// @Success 200 {string} string
// @Router /api/v1/token/refresh [post]
func (h *Handler) Refresh(c *gin.Context) {

}
