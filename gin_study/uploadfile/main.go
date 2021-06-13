package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 给表单限制上传大小
	//engine.MaxMultipartMemory = 8 << 20  // 8M
	// 2. 创建路由规则
	engine.POST("/upload", func(c *gin.Context) {
		// 从表单获取文件
		file, _ := c.FormFile("file")

		// 日志记录一下当前上传的文件名
		fmt.Println(file.Filename)

		// 将当前上传的文件保存在某一个指定目录下(默认在当前目录)
		// 上传文件的文件名可以由用户自定义，所以可能包含非法字符串，为了安全起见，应该由服务端统一文件名规则
		dst := "./" + file.Filename
		_ = c.SaveUploadedFile(file, dst)

		// 响应上传成功
		c.String(http.StatusOK, fmt.Sprintf("%v 上传成功！", file.Filename))
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
