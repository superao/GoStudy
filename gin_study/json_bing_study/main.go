package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义绑定数据的结构体
type login struct {
	User     string `json:"username" form:"user" uri:"user" xml:"user" binding:"required"`
	Password string `json:"password" form:"password" uri:"password" xml:"password" binding:"required"`
}

// 客户端采用 json 格式传参，服务器解析参数并绑定在结构体上
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建指定路由的处理方法
	engine.POST("/login", func(c *gin.Context) {
		// 声明绑定数据的变量
		var json login

		// 将 request 的 body 中的数据，自动按照 json 格式解析到结构体
		err := c.ShouldBindJSON(&json)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// 检测用户名和密码是否正确
		if json.User != "root" || json.Password != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"status": "304"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	// 3. 启动路由引擎
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}

// 测试与结果:
//C:\Users\www18>curl http://127.0.0.1:8000/login -X POST -H "content-type: application/json" -d "{\"username\": \"root\", \"password\": \"admin\"}"
//{"status":"200"}
//C:\Users\www18>curl http://127.0.0.1:8000/login -X POST -H "content-type: application/json" -d "{\"username\": \"root\", \"password\": \"123456\"}"
//{"status":"304"}
