package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 本示例演示 Gin 框架的重定向
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	engine.GET("/baidu", func(c *gin.Context) {
		// Gin 框架支持内部重定向和外部重定向(外部重定向)
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
	})

	engine.GET("/index", func(c *gin.Context) {
		// Gin 框架支持内部重定向和外部重定向(内部重定向)
		c.Redirect(http.StatusMovedPermanently, "/hello")
	})

	engine.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world!")
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
