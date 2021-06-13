package main

import (
	"fmt"
	"net/rpc"
)

// Params 定义 RPC 客户端消息格式
type Params struct {
	Width  int
	Height int
}

// RPC 客户端的两个流程
// 1. 连接远程的 RPC 服务器
// 2. 调用远程 RPC 服务器提供的方法

func main() {
	// 连接远程的 RPC 服务器
	client, err := rpc.DialHTTP("tcp", ":8000")
	if err != nil {
		fmt.Println("连接远程 RPC 服务器失败！")
		return
	}

	// 调用远程的 RPC 方法
	var ret int
	// Call: 远程的服务名、服务参数、服务返回值
	err = client.Call("Rect.Area", Params{Width: 10, Height: 20}, &ret)
	if err != nil {
		fmt.Println("远程服务调用失败！")
		return
	}
	fmt.Println("矩形的面积为: ", ret)

	err = client.Call("Rect.Perimeter", Params{Width: 10, Height: 5}, &ret)
	if err != nil {
		fmt.Println("远程服务调用失败！")
		return
	}
	fmt.Println("矩形的周长为: ", ret)
}
