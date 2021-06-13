package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// Params 定义 RPC 的请求参数结构
type Params struct {
	Width  int
	Height int
}

// Rect 定义服务端的服务对象类型
// 一般都是将相关服务的函数绑定在某一个类型上
type Rect struct {
}

// Area 定义 RPC 服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// Perimeter 定义 RPC 服务端方法，求矩形周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Height + p.Width) * 2
	return nil
}

// RPC 服务端的三个流程:
// 1. 注册服务
// 2. 将服务绑定在某一个通信协议上，如: http
// 3. 监听服务

func main() {
	rect := new(Rect)

	// 1. 注册服务
	rpc.Register(rect)

	// 2. 将服务绑定在 http 通信协议上
	rpc.HandleHTTP()

	// 3. 启动 RPC 服务器进行监听
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("RPC 服务器启动失败！")
		return
	}
}
