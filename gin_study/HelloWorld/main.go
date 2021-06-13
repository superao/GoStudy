package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1. 创建路由引擎
	// Default 默认使用了两个中间件 Logger(), Recovery()
	// 当然在使用的过程中，也是可以创建没有使用任何中间件的路由引擎
	// engine := gin.New()
	engine := gin.Default()

	// 2. 在路由引擎上绑定路由规则，以及对应的执行函数
	// gin.Context 封装了原有 http 库的 request 和 response 库
	// 使得原有对 http 的相关操作更加的快捷，简单
	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})
	// engine.POST("/somePath", posting)
	// engine.PUT("/somePath", putting)
	// engine.DELETE("/somePath", deleting)
	// engine.HEAD("/somePath", head)
	// engine.OPTIONS("/somePath", options)

	// 3. 路由引擎监听指定端口，默认为 8080
	err := engine.Run("127.0.0.1:8000")
	if err != nil {
		fmt.Println("路由引擎运行失败")
	}
}
