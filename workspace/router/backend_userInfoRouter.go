package Router

import (
	"fmt"

	"Middleware"
	"controller"

	"github.com/gin-gonic/gin"
)

func InitUserInfoRouter(r *gin.Engine) {
	fmt.Println("InitUserInfoRouter")
	user := r.Group("/geeker/users")
	// user.POST("/login", func(c *gin.Context) {

	// })
	//查询用户信息
	user.GET("/infoList", Middleware.AuthMiddleWare(1), func(c *gin.Context) {

		controller.UserController{}.GetUserInfoList(c)
	})
	user.DELETE("/infoList", Middleware.AuthMiddleWare(1), func(c *gin.Context) {
		controller.UserController{}.DeleteUser(c)
	})
	//更新
	user.PUT("/infoList", Middleware.AuthMiddleWare(1), func(c *gin.Context) {
		controller.UserController{}.UpdateUser(c)
	})
	// Code to initialize router
}
