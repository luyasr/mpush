package response

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/pkg/errorhandler"
	"net/http"
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func JSON(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Data: data, Message: "success"})
}

func JSONWithError(c *gin.Context, err error) {
	defer c.Abort()

	httpCode, bizCode, message := errorhandler.ErrorInfo(err)
	c.JSON(httpCode, Response{Code: bizCode, Data: nil, Message: message})
}
