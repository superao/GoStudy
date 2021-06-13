package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由组
	// 路由组下面的大括号并不是语法，只不过这样写是一种规范
	v1 := engine.Group("/v1")
	{
		v1.GET("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "hello world!")
		})

		v1.GET("/hi", func(c *gin.Context) {
			c.String(http.StatusOK, "hi~ kukizhang")
		})
	}

	v2 := engine.Group("/v2")
	{
		v2.POST("/hello", func(c *gin.Context) {
			c.String(http.StatusOK, "HELLO WORLD!")
		})

		v2.POST("/hi", func(c *gin.Context) {
			c.String(http.StatusOK, "HI~ KUKIZHANG")
		})
	}

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}

// 测试与结果:
//C:\Users\www18>curl http://127.0.0.1:8000/v1/hello
//hello world!
//C:\Users\www18>curl http://127.0.0.1:8000/v1/hi
//hi~ kukizhang
//C:\Users\www18>curl http://127.0.0.1:8000/v2/hi -X POST
//HI~ KUKIZHANG
//C:\Users\www18>curl http://127.0.0.1:8000/v2/hello -X POST
//HELLO WORLD!
