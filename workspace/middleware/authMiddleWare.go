package Middleware

import (
	"config"
	"strings"
	"utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleWare(role int) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err == "1" {
					c.JSON(200, gin.H{"code": 404, "msg": "权限不足"})
				} else {

					c.JSON(200, gin.H{"code": 401, "msg": "token invalid"})

				}
				c.Abort()
				return
			}
		}()
		//获取 authorization Barear token
		auth := c.GetHeader("Authorization")

		parts := strings.Split(auth, " ")

		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			claims, err := utils.ParseToken(token, config.ConfigList.Jwt.Secret)
			if err != nil {
				panic("")
			}
			//判断角色
			if int(claims["role"].(float64)) == role {
				//将用户信息传递下去
				c.Set("role", claims["role"])
				c.Next()
			} else {
				panic("1")
			}
		} else {
			panic("")
		}

	}
}
