package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"task4/internal/util"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 是否有token
		auth := c.GetHeader("Authorization")
		if auth == "" {
			util.Unauthorized(c, "缺少认证Token")
			c.Abort()
			return
		}

		// 2. token格式检验
		n := strings.SplitN(auth, " ", 2)
		if !(len(n) == 2 && n[0] == "Bearer") {
			util.Unauthorized(c, "Token格式错误")
		}

		// 3. token数据校验
		token := n[1]
		claims, err := util.ParseToken(token)
		if err != nil {
			util.Unauthorized(c, "Invalid token")
			c.Abort()
			return
		}

		// 将用户信息写入上下文
		c.Set("user_id", claims.ID)
		c.Set("user_name", claims.Username)

		c.Next()
	}
}
