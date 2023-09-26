package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"xzs/common"
	"xzs/pkg/jwtutil"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := checkAuth(c)
		if err != nil {
			common.ResponseNoLogin(c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func checkAuth(c *gin.Context) (err error) {
	session := sessions.Default(c)
	if session.Get("userName") == nil {
		// 这里还需要判断是否记住密码的逻辑，将cookie中的userName设置到session中
		token, err := c.Cookie("token")
		if err != nil {
			return err
		}
		// 解析token
		claims, err := jwtutil.ParseToken(token)
		if err != nil {
			return err
		}
		userName := claims.UserName
		// 设置session
		session.Set("userName", userName)
		session.Save()
		return err
	} else {
		// 将userName存入context
		c.Set("userName", session.Get("userName"))
	}
	return
}
