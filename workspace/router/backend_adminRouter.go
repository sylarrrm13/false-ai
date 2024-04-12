package Router

import (
	"Middleware"
	"controller"

	"github.com/gin-gonic/gin"
)

// type loginData struct {
// 	Accesstoken  string `json:"access_token"`
// 	RefreshToken string `json:"refresh_token"`
// }

// type loginRes struct {
// 	Code int       `json:"code"`
// 	Data loginData `json:"data"`
// }

func InitAdminRouter(r *gin.Engine) {
	admin := r.Group("/geeker")
	admin.POST("/login", func(c *gin.Context) {
		controller.AdminController{}.Login(c)
	})
	admin.GET("/admins", func(c *gin.Context) {
		controller.AdminController{}.GetAdmins(c)

	})
	admin.PUT("/adminInfo", Middleware.AuthMiddleWare(1), func(ctx *gin.Context) {
		//更新管理员密码，判断原密码是否正确
		controller.AdminController{}.UpdatePassword(ctx)

	})

	//验证refresh token
	admin.POST("/refreshToken", func(c *gin.Context) {
		controller.AdminController{}.RefreshToken(c)
	})

	// Code to initialize router
}
