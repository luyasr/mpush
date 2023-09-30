package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/common/response"
	"net/http"
	"strconv"
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
		group.DELETE(":id", h.DeleteUser)
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}
	user, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}

	h.service.log.Info().Msg(fmt.Sprintf("创建用户%s成功", user.Username))
	c.JSON(http.StatusOK, response.New(user))
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	parseInt64, _ := strconv.ParseInt(id, 10, 64)
	err := h.service.DeleteById(c.Request.Context(), parseInt64)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}
	h.service.log.Info().Msg(fmt.Sprintf("删除用户id%d成功", parseInt64))
	c.JSON(http.StatusOK, response.New(nil))
}
