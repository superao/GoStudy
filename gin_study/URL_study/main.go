package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取 URL 中问号后面的多个 KV 键值对参数
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	engine.GET("/hello", func(c *gin.Context) {
		value1 := c.DefaultQuery("name", "jack") // 若为空，则有默认值
		value2 := c.Query("id")                  // 若为空，则为空串

		c.JSON(http.StatusOK, gin.H{"name": value1, "id": value2})
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}

// 测试域名
// http://127.0.0.1:8000/hello?name=zhangsan&id=123
// 结果: {"id":"123","name":"zhangsan"}
