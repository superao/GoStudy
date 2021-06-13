package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取 URL 中的 API 参数
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建路由规则
	// 此规则能够匹配到 /user/zhangsan 这种格式，但是不能匹配到 /user/ 和 /user 以及 /user/zhangsan/xxx 这种格式
	//engine.GET("/user/:name", func(c *gin.Context) {
	//	// 获取 API 参数
	//	str := c.Param("name")
	//	c.String(http.StatusOK, str)
	//})

	// 此规则能够匹配到 /user/zhangsan/、/user/zhangsan/lisi 这种格式
	engine.GET("/user/:name/*action", func(c *gin.Context) {
		// 获取 API 参数
		str1 := c.Param("name")
		str2 := c.Param("action")
		c.String(http.StatusOK, str1+" and "+str2)
	})

	// 3. 监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启动失败！")
		return
	}
}

// 匹配规则总结：
// 1. /:name 冒号和占位符在一起相当于在这个 URL 中这个占位符的位置必须有参数，如果没有参数那么就会匹配不到，
// 并且就算该位置有参数匹配到了，但是后续依然不能有参数，如果后续依然有参数，那么还是会匹配不到。
// 2. /*name 星号和占位符在一起相当于在这个 URL 中这个占位符的位置的参数可有可无，它都会匹配得到。
