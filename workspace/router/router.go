package Router

import (
	"Middleware"

	"github.com/gin-gonic/gin"
)

// 部署的时候删掉

func InitRouter(r *gin.Engine) {
	r.Use(Middleware.CorsMiddleware())
	InitAdminRouter(r)
	InitUserInfoRouter(r)
	InitConfigRouter(r)
	InitModelCate(r)
	InitModelKey(r)
	InitCardTypeRouter(r)
	InitFrontUser(r)
}
