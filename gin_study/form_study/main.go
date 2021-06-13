package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取前端 form 表单中的各种数据
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	engine.POST("/form", func(c *gin.Context) {
		// 获取 form 表单中的各个参数
		t := c.DefaultPostForm("type", "kukizhang") // 表单中没该参数，故结果为kukizhang

		username := c.PostForm("username")
		password := c.PostForm("password")
		hobby := c.PostForm("hobby")

		// 返回获取到的参数结果
		c.JSON(http.StatusOK, gin.H{"username": username, "type": t,
			"password": password, "hobby": hobby})
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
