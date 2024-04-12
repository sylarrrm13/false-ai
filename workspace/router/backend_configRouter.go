package Router

import (
	"fmt"

	"controller"

	"github.com/gin-gonic/gin"
)

func InitConfigRouter(r *gin.Engine) {
	fmt.Println("InitConfigRouter")
	config := r.Group("/geeker/config/:listname")

	//查询用户信息
	config.GET("", func(c *gin.Context) {

		controller.ConfigController{}.GetConfig(c)
	})
	config.PUT("", func(c *gin.Context) {
		controller.ConfigController{}.UpdateConfig(c)

	})
	// config.PUT("/email", func(c *gin.Context) {
	// 	configController.ConfigController{}.UpdateConfig(c, "Email")
	// })
	// Code to initialize router
}
