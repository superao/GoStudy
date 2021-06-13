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
		// 从表单中获取多个表单数据
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, "获取表单数据失败！")
			fmt.Println("获取表单数据失败！")
			return
		}

		// 从表单数据中获取指定字段的多个文件
		files := form.File["files"]

		// 遍历每一个文件，然后进行保存
		for _, file := range files {
			err := c.SaveUploadedFile(file, file.Filename)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("%v 文件上传失败！", file.Filename))
				fmt.Println("文件储存失败！")
				return
			}
		}

		// 向客户端响应结果
		c.String(http.StatusOK, fmt.Sprintf("%v 个文件上传成功！", len(files)))
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}
