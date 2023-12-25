package middleware

import (
	"blog/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware 用于验证tokens
func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// cookie验证jwtToken
		tokenString, err := c.Cookie("jwtToken")
		// 获取token错误或token失效
		if err != nil || controller.RevokedTokens[tokenString] {
			fmt.Println("BadReq")
			//utils.RespondWithError(401, "Unauthorized", c)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		mc, err := controller.ParseToken(tokenString)
		if err != nil {
			fmt.Println("token error")
			//utils.RespondWithError(401, "Unauthorized", c)
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			return
		}
		c.Set("user", mc.Username)
		c.Next()
	}
}
