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
		group.PUT(":id", h.UpdateUser)
		group.GET(":id", h.GetUserById)
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

func (h *Handler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	err := c.BindJSON(&req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}
	id := c.Param("id")
	parseInt64, _ := strconv.ParseInt(id, 10, 64)

	err = h.service.Update(c.Request.Context(), parseInt64, &req)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}
	h.service.log.Info().Msg(fmt.Sprintf("用户id%d更新成功", parseInt64))
	c.JSON(http.StatusOK, response.New(nil))
}

func (h *Handler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	parseInt64, _ := strconv.ParseInt(id, 10, 64)

	byId, err := h.service.GetById(c.Request.Context(), parseInt64)
	if err != nil {
		h.service.log.Error().Stack().Err(err).Send()
		c.JSON(http.StatusOK, response.NewWithError(err))
		return
	}

	h.service.log.Info().Msg(fmt.Sprintf("用户id%d查询", parseInt64))
	c.JSON(http.StatusOK, response.New(byId))
}
