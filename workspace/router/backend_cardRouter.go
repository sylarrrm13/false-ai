package Router

import (
	"controller"
	"fmt"

	"Middleware"

	"github.com/gin-gonic/gin"
)

//新增卡类型的router

func InitCardTypeRouter(r *gin.Engine) {

	card := r.Group("/geeker/card")
	card.Use(Middleware.AuthMiddleWare(1))
	card.GET("", func(c *gin.Context) {
		controller.CardController{}.GetCardList(c)
	})
	card.POST("", func(c *gin.Context) {
		controller.CardController{}.GenCard(c)

	})
	card.POST("/export/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		// 设置响应头，告诉客户端这是一个文件
		c.Header("Content-Disposition", "attachment; filename="+filename)
		c.Header("Content-Type", "application/octet-stream")

		// 将文件内容作为响应的正文发送给客户端
		c.File("../files/card/" + filename)
	})

	card.POST("/", func(c *gin.Context) {
		controller.CardController{}.GenCard(c)

	})

	card.DELETE("/", func(c *gin.Context) {
		controller.CardController{}.DeleteCard(c)

	})
	card.DELETE("", func(c *gin.Context) {
		controller.CardController{}.DeleteCard(c)

	})

	//cardType部分

	cardType := r.Group("/geeker/cardType")
	cardType.Use(Middleware.AuthMiddleWare(1))

	cardType.GET("", func(c *gin.Context) {
		fmt.Println("get cardType")
		controller.CardTypeController{}.GetCardTypeList(c)

	})

	cardType.GET("/", func(c *gin.Context) {
		fmt.Println("get cardType")
		controller.CardTypeController{}.GetCardTypeList(c)

	})

	cardType.POST("/", func(c *gin.Context) {
		controller.CardTypeController{}.AddCardType(c)
	})

	cardType.PUT("/", func(c *gin.Context) {
		controller.CardTypeController{}.UpdateCardType(c)
	})

	cardType.DELETE("/", func(c *gin.Context) {
		controller.CardTypeController{}.DeleteCardType(c)
	})

	// Code to initialize router
}
