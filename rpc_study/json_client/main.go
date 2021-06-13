package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// Params 定义 RPC 客户端请求参数格式
type Params struct {
	Num1 int
	Num2 int
}

func main() {
	// 1. 连接远程 RPC 服务器
	// cli, err := jsonrpc.Dial("tcp", ":8000")
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("连接远程 RPC 服务器失败！")
		return
	}
	cli := jsonrpc.NewClient(conn)

	// 2. 调用远程 RPC 服务器提供的服务
	// 加法服务
	var ret int
	err = cli.Call("Response.Sum", Params{Num1: 1, Num2: 2}, &ret)
	if err != nil {
		fmt.Println("远程服务调用失败！")
		return
	}
	fmt.Println("求和结果为: ", ret)

	// 减法服务
	err = cli.Call("Response.Diminish", Params{Num1: 1, Num2: 2}, &ret)
	if err != nil {
		fmt.Println("远程服务调用失败！")
		return
	}
	fmt.Println("求差结果为: ", ret)
}

