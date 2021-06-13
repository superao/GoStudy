package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 本示例展示同步处理请求和异步处理请求的区别
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则对应的处理方法
	// 异步处理请求
	engine.GET("/async", func(c *gin.Context) {
		// 拷贝出一个上下文副本
		cp := c.Copy()

		// 创建 goroutine 进行异步处理，然后此时接受请求的过程先结束
		go func() {
			time.Sleep(3 * time.Second)
			fmt.Println("异步处理请求：", cp.Request.URL.String())
		}()
		// 异步请求不阻塞当前操作，请求的处理由另一个 goroutine 去处理

		// 直接向客户端响应请求处理成功
		c.String(http.StatusOK, "异步: 请求处理成功！")
	})

	// 同步处理请求
	engine.GET("/sync", func(c *gin.Context) {
		// 当前线程同步处理请求
		time.Sleep(3 * time.Second)
		fmt.Println("异步处理请求：", c.Request.URL.String())

		// 向客户端响应请求处理成功
		c.String(http.StatusOK, "同步: 请求处理成功！")
	})

	// 3. 启动路由引擎
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
