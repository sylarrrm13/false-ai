package Middleware

import (
	"config"
	"strings"
	"utils"

	"github.com/gin-gonic/gin"
)

func UserAuthMiddleWare(authType int) gin.HandlerFunc {
	//如果auType =1，先判断 Authorization是否存在
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if authType == 1 {
			//判断auth是否存在
			if auth == "" {
				c.Set("userState", "0")
			} else {
				parts := strings.Split(auth, " ")
				if len(parts) == 2 && parts[0] == "Bearer" {
					token := parts[1]
					claims, err := utils.ParseToken(token, config.ConfigList.Jwt.Secret)
					if err != nil {
						c.Set("userState", "0")
					} else {
						//判断 role是否为3
						if int(claims["role"].(float64)) != 3 {
							c.Set("userState", "0")
						} else {
							c.Set("userState", "1")
							c.Set("userId", claims["userId"])
						}

					}
				} else {
					c.Set("userState", "0")
				}
			}
		} else {
			if auth == "" {
				//直接返回未授权
				c.JSON(401, gin.H{"code": 401, "msg": "授权到期，请重新登录"})
				c.Abort()
				return
			}

			parts := strings.Split(auth, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				c.JSON(401, gin.H{"code": 401, "msg": "授权到期，请重新登录"})
				c.Abort()
				return
			} else {
				token := parts[1]
				claims, err := utils.ParseToken(token, config.ConfigList.Jwt.Secret)
				if err != nil {
					c.JSON(401, gin.H{"code": 401, "msg": "授权到期，请重新登录"})
					c.Abort()
					return
				} else {
					//判断 role是否为3
					if int(claims["role"].(float64)) != 3 {
						c.JSON(401, gin.H{"code": 401, "msg": "授权到期，请重新登录"})
						c.Abort()
						return
					} else {
						c.Set("userId", claims["userId"])
					}
				}
			}

		}

	}

}
