package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/luyasr/gaia/ioc"
	"github.com/luyasr/gaia/transport/http/response"
	"github.com/luyasr/mpush/apps/token"
	"github.com/luyasr/mpush/apps/user"
)

type Auth struct {
	token token.Service
	Role  user.Role
}

func NewAuth() *Auth {
	return &Auth{
		token: ioc.Container.Get(ioc.ControllerNamespace, token.Name).(token.Service),
	}
}

func (a *Auth) Token(c *gin.Context) {
	// 从请求头中获取token
	tk := c.Request.Header.Get(token.Name)
	// 校验token
	validatedToken, err := a.token.Validate(c, tk)
	if err != nil {
		response.GinJsonWithError(c, err)
		return
	}

	// 将鉴权后的token放入上下文
	ctx := context.WithValue(c.Request.Context(), token.Name, validatedToken)
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}
