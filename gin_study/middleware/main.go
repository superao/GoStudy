package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Middleware 定义中间件(一般中间件用于单点登陆)
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了！")
		// 在中间件中设置变量到 Context 中，后续可以通过 Get 进行获取
		c.Set("request", "中间件")

		// 执行后续服务
		c.Next()

		// 服务执行完毕后，在中间件中进行后续处理
		// 获取服务响应状态码
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕，结果为: ", status)
		t2 := time.Since(t)
		fmt.Println("整个流程执行的时间为: ", t2)
	}
}

// 本示例演示中间件的使用 (局部中间件、全局中间件)
func main() {
	// 1. 定义路由引擎
	engine := gin.Default()

	// 2. 注册/使用中间件
	engine.Use(Middleware())
	// {} 是为了代码规范，表示括号里面的服务都使用了当前的中间件
	{
		engine.GET("/middleware", func(c *gin.Context) {
			// 获取中间件中所设置的值
			value, _ := c.Get("request")
			fmt.Println("request: ", value)

			// 向客户端响应数据
			c.JSON(http.StatusOK, gin.H{"request": value})
		})

		// 局部中间件的使用
		// 当访问该路由的时候，上述定义的中间件会执行两遍
		engine.GET("/middleware2", Middleware(), func(c *gin.Context) {
			// 获取中间件中所设置的值
			value, _ := c.Get("request")
			fmt.Println("request: ", value)

			// 向客户端响应数据
			c.JSON(http.StatusOK, gin.H{"request": value})
		})
	}

	// 3. 启动路由引擎
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
