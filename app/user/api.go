package user

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/pkg/response"
	"github.com/luyasr/mpush/pkg/utils"
)

type Handler struct {
	service *Service
}

func NewHandler() *Handler {
	return &Handler{
		service: NewService(),
	}
}

func (h *Handler) Registry(r *gin.RouterGroup) {
	group := r.Group("user")
	{
		group.POST("", h.CreateUser)
		group.DELETE(":id", h.DeleteUserByID)
		group.PUT(":id", h.UpdateUserByID)
		group.GET(":id", h.QueryUserByID)
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req *CreateUserRequest

	err := c.BindJSON(&req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("创建用户%s失败", req.Username)
		response.JSONWithError(c, err)
		return
	}

	user, err := h.service.CreateUser(c.Request.Context(), req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("创建用户%s失败", req.Username)
		response.JSONWithError(c, err)
		return
	}

	h.service.log.Info().Msgf("创建用户%s成功", user.Username)
	response.JSON(c, user)
}

func (h *Handler) DeleteUserByID(c *gin.Context) {
	var req *DeleteUserRequest

	req.ID = utils.StringToInt64(c.Param("id"))

	err := h.service.DeleteUser(c.Request.Context(), req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("删除用户id%d失败", req.ID)
		response.JSONWithError(c, err)
		return
	}

	h.service.log.Info().Msgf("删除用户id%d成功", req.ID)
	response.JSON(c, nil)
}

func (h *Handler) UpdateUserByID(c *gin.Context) {
	var req *UpdateUserRequest

	err := c.BindJSON(req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("更新用户id%d失败", req.ID)
		response.JSONWithError(c, err)
		return
	}

	req.ID = utils.StringToInt64(c.Param("id"))

	err = h.service.UpdateUser(c.Request.Context(), req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("更新用户id%d失败", req.ID)
		response.JSONWithError(c, err)
		return
	}
	h.service.log.Info().Msgf("更新用户id%d成功", req.ID)
	response.JSON(c, nil)
}

func (h *Handler) QueryUserByID(c *gin.Context) {
	req := NewQueryUserByIdRequest(utils.StringToInt64(c.Param("id")))

	user, err := h.service.QueryUser(c.Request.Context(), req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Msgf("查询用户id%v失败", req.QueryByValue...)
		response.JSONWithError(c, err)
		return
	}

	h.service.log.Info().Msgf("查询用户id%v成功", req.QueryByValue...)
	response.JSON(c, user)
}
