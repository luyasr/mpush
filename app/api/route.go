package api

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/common/response"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, response.New("pong"))
}
