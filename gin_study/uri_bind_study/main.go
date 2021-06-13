package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义绑定数据的结构体
type login struct {
	User     string `uri:"user"`
	Password string `uri:"password"`
}

// 客户端采用 uri 格式传参，服务器解析参数并绑定在指定结构体上
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	engine.GET("/:user/:password", func(c *gin.Context) {
		// 声明接受数据的变量
		var uri login

		// 将 rui 中的数据，自动解析并绑定到结构体上
		err := c.ShouldBindUri(&uri)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// 检测用户名和密码是否正确
		if uri.User != "root" || uri.Password != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"status": "403"})
			return
		}

		// 向客户端响应结果
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// 3. 启动路由引擎
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
