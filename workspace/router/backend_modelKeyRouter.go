package Router

import (
	"controller"

	"github.com/gin-gonic/gin"
)

func InitModelKey(r *gin.Engine) {
	modelKey := r.Group("/geeker/modelKey")

	//查询用户信息
	modelKey.GET("", func(c *gin.Context) {
		controller.ModelKeyController{}.GetModelKeyList(c)
	})
	modelKey.PUT("", func(c *gin.Context) {
		controller.ModelKeyController{}.UpdateModelKey(c)

	})
	modelKey.PUT("/", func(c *gin.Context) {
		controller.ModelKeyController{}.UpdateModelKey(c)

	})
	//新增模型类
	modelKey.POST("/", func(c *gin.Context) {
		controller.ModelKeyController{}.AddModelKey(c)
	})
	modelKey.POST("", func(c *gin.Context) {
		controller.ModelKeyController{}.AddModelKey(c)
	})

	//删除模型类
	modelKey.DELETE("", func(c *gin.Context) {
		controller.ModelKeyController{}.DeleteModelKey(c)
	})
	modelKey.DELETE("/", func(c *gin.Context) {
		controller.ModelKeyController{}.DeleteModelKey(c)
	})
	// config.PUT("/email", func(c *gin.Context) {
	// 	configController.ConfigController{}.UpdateConfig(c, "Email")
	// })
	// Code to initialize router
}
