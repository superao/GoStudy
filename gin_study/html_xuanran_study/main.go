package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 本示例演示后端如何渲染 HTML 页面
func main() {
	// 1. 创建路由
	engine := gin.Default()

	// 2. 加载指定的模板文件
	engine.LoadHTMLGlob("templates/*") // 加载 templates 目录下所有的文件
	//engine.LoadHTMLFiles("templates/index.tmpl")   // 加载指定的文件

	// 3. 创建指定的路由规则
	engine.GET("/index", func(c *gin.Context) {
		// 这里我们根据指定的 html 模板文件进行渲染，本质上就是将字符串进行一下替换即可
		// 本质就是将 tmpl 模板文件中的 title 替换成了 hello world!
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"title":"hello world!"})
	})

	// 4. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
