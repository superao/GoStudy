package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义绑定数据的结构体
type login struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
}

// 客户端采用 xml 格式传数据，服务器解析数据并绑定在结构体上
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	engine.POST("/login", func(c *gin.Context) {
		// 声明接受数据的变量
		var xml login

		// 将 request 的 body 中的数据，自动按照 xml 格式解析到结构体
		err := c.ShouldBindXML(&xml)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// 检测用户名和密码是否正确
		if xml.User != "root" || xml.Password != "admin" {
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

// xml 格式数据：application/xml
//<?xml version="1.0" encoding="utf-8"?>
//
//<root>
//<user>root</user>
//<password>admin</password>
//</root>
