package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

// 本示例学习 gin 框架中各种数据格式的渲染，即各种数据类型的返回
func main() {
	// 1. 创建路由引擎
	engine := gin.Default()

	// 2. 创建理由规则
	// json 格式返回
	engine.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "someJSON", "status": 200})
	})

	// 结构体数据返回
	engine.GET("/someStruct", func(c *gin.Context) {
		// 创建需要返回的结构体变量
		var person struct {
			Name string
			Id   uint64
		}

		// 为结构体赋值
		person.Name = "kukizhang"
		person.Id = 123456

		// 向客户端以 JSON 格式响应数据
		c.JSON(http.StatusOK, gin.H{"name": person.Name, "id": person.Id})
	})

	// XML 格式返回
	engine.GET("/somexml", func(c *gin.Context) {
		// 服务端以 json 格式组织数据，然后 gin 框架会自动的将我们想要返回的数据转化为 xml 格式，最后响应给客户端
		c.XML(http.StatusOK, gin.H{"name": "kukizhang", "id": "123456"})
	})

	// yaml 格式返回
	// yaml 格式的数据经常作为配置文件格式进行使用，当我们访问这个格式的数据时，它会自动的下载到本地
	engine.GET("/someyaml", func(c *gin.Context) {
		// 文本格式
		c.YAML(http.StatusOK, gin.H{"name": "kukizhang", "id": "123456"})
	})

	// protobuf 格式返回
	// pb 格式是一种非常高效的用于数据传输的格式，当我们访问这个格式的数据时，它会自动的下载到本地
	engine.GET("/somepb", func(c *gin.Context) {
		// 二进制格式
		// 定义数据(下文为数据格式)
		labal := "test"
		reps := []int64{int64(123), int64(456)}

		// 序列化指定数据
		data := &protoexample.Test{
			Label: &labal,
			Reps:  reps,
		}

		// 向客户端响应对应的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	// 3. 启动路由引擎，监听指定端口
	err := engine.Run(":8000")
	if err != nil {
		fmt.Println("路由引擎启失败！")
		return
	}
}

//type Test struct {
//	Label            *string             `protobuf:"bytes,1,req,name=label" json:"label,omitempty"`
//	Type             *int32              `protobuf:"varint,2,opt,name=type,def=77" json:"type,omitempty"`
//	Reps             []int64             `protobuf:"varint,3,rep,name=reps" json:"reps,omitempty"`
//	Optionalgroup    *Test_OptionalGroup `protobuf:"group,4,opt,name=OptionalGroup" json:"optionalgroup,omitempty"`
//	XXX_unrecognized []byte              `json:"-"`
//}
