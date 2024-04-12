package Router

import (
	"controller"

	"github.com/gin-gonic/gin"
)

func InitModelCate(r *gin.Engine) {
	modelCate := r.Group("/geeker/modelCate")

	//查询用户信息
	modelCate.GET("", func(c *gin.Context) {
		controller.ModelCateController{}.GetModelCateList(c)
	})
	modelCate.PUT("", func(c *gin.Context) {
		controller.ModelCateController{}.UpdateModelCate(c)

	})
	modelCate.PUT("/", func(c *gin.Context) {
		controller.ModelCateController{}.UpdateModelCate(c)

	})
	//新增模型类
	modelCate.POST("/", func(c *gin.Context) {
		controller.ModelCateController{}.AddModelCate(c)
	})
	modelCate.POST("", func(c *gin.Context) {
		controller.ModelCateController{}.AddModelCate(c)
	})

	//删除模型类
	modelCate.DELETE("", func(c *gin.Context) {
		controller.ModelCateController{}.DeleteModelCate(c)
	})
	modelCate.DELETE("/", func(c *gin.Context) {
		controller.ModelCateController{}.DeleteModelCate(c)
	})
	// config.PUT("/email", func(c *gin.Context) {
	// 	configController.ConfigController{}.UpdateConfig(c, "Email")
	// })
	// Code to initialize router
}
