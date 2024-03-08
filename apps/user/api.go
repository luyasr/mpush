package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/transport/http/response"
)

type Handler struct {
	service Service
}

func init() {
	ioc.Container.Registry(ioc.HandlerNamespace, &Handler{})
}

func (h *Handler) Init() error {
	h.service = ioc.Container.Get(ioc.ControllerNamespace, Name).(Service)

	return nil
}

func (h *Handler) Name() string {
	return Name
}

func (h *Handler) Registry(r gin.IRouter) {
	r = r.Group("/user")
	{
		r.POST(".", h.Create)
		r.GET(":id", h.Query)
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

	user, err := h.service.Create(c, req)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, user)
}

// Query 查询用户
// @Summary 查询用户
// @Description 查询用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} User
// @Router /api/v1/user/{id} [get]
func (h *Handler) Query(c *gin.Context) {
	parseInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}
	user, err := h.service.QueryById(c, parseInt)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, user)
}

// Delete 删除用户
// @Summary 删除用户
// @Description 删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200
// @Router /api/v1/user/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	parseInt, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	if err = h.service.Delete(c, parseInt); err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	response.GinJson(c, nil)
}
