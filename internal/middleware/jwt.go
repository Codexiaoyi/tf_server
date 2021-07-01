package middleware

import (
	"strings"
	"tfserver/internal/errmsg"
	"tfserver/internal/response"
	"tfserver/pkg/jwt"

	"github.com/gin-gonic/gin"
)

//基于JWT的认证中间件
func JwtMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		//使用Bearer认证方式
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			response.Response(c, errmsg.TOKEN_NOT_FOUND)
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			response.Response(c, errmsg.TOKEN_FORMAT_ERROR)
			c.Abort()
			return
		}

		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			response.Response(c, errmsg.TOKEN_NOT_VALID)
			c.Abort()
			return
		}
		// 将当前请求的email信息保存到请求的上下文c上
		c.Set("email", mc.Email)
		c.Next() // 后续的处理函数可以用过c.Get("email")来获取当前请求的用户信息
	}
}
