package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/luyasr/mpush/config"
	"github.com/luyasr/mpush/pkg/jwt"
	"github.com/luyasr/mpush/pkg/response"
	"golang.org/x/net/context"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(jwt.Name)
		claims, err := jwt.ParseJwt(token, config.C.Jwt.Secret)
		if err != nil {
			response.JSONWithError(c, err)
		}

		ctx := context.WithValue(c.Request.Context(), "username", claims.Username)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
