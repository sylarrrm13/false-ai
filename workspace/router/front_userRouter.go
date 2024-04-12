package Router

import (
	"Middleware"
	"controller"

	"github.com/gin-gonic/gin"
)

func InitFrontUser(r *gin.Engine) {
	frontUser := r.Group("/ftapi")
	frontUser.GET("/initData", Middleware.UserAuthMiddleWare(1), func(c *gin.Context) {
		controller.FrontUserController{}.InitData(c)
	})
	frontUser.POST("/sendEmailCode", func(c *gin.Context) {
		controller.FrontUserController{}.SendEmailCode(c)
	})
	frontUser.POST("/register", func(c *gin.Context) {
		controller.FrontUserController{}.Register(c)
	})
	frontUser.POST("/login", func(c *gin.Context) {
		controller.FrontUserController{}.Login(c)
	})
	//获取用户Bill信息
	frontUser.GET("/bill_info", Middleware.UserAuthMiddleWare(2), func(c *gin.Context) {
		controller.FrontUserController{}.BillInfo(c)
	})

	frontUser.POST("/v1/chat/completions", Middleware.UserAuthMiddleWare(2), func(c *gin.Context) {
		controller.FrontUserController{}.Chat(c)

	})

	//RECHARGE
	frontUser.POST("/recharge", Middleware.UserAuthMiddleWare(2), func(c *gin.Context) {
		controller.FrontUserController{}.Recharge(c)
	})

	frontUser.POST("/pre_sign", Middleware.UserAuthMiddleWare(2), func(c *gin.Context) {
		controller.FrontUserController{}.PreSign(c)
	})
}
