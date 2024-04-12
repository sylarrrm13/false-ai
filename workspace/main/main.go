package main

import (
	"Router"
	_ "config"

	"github.com/gin-gonic/gin"
)

func main() {

	// 读取 config.ini 配置文件

	r := gin.Default()
	// logger := &lumberjack.Logger{
	// 	Filename:   "./log/gin.log",

	// 	MaxSize:    500, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     1, //days
	// }

	// gin.DefaultWriter = logger
	// logrus.SetOutput(gin.DefaultWriter)
	Router.InitRouter(r)
	r.Run("192.168.1.5:3200")

}
