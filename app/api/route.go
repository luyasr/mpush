package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/pkg/response"
)

func Ping(c *gin.Context) {
	response.JSON(c, "pong")
}
