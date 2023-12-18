package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	hello := "Hello World!"
	//1.创建路由
	r := gin.Default()
	//2.绑定路由规则，执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, hello)
	})

	r.GET("/ping", func(c * gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})
	//3.监听端口，默认8080w
	r.Run(":10001")
}
