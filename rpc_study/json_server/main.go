package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// Params 定义 RPC 服务端接受参数的格式
type Params struct {
	Num1 int
	Num2 int
}

// Response 定义 RPC 服务端所能提供的各种服务
type Response struct {
}

// Sum 求和服务
func (r *Response) Sum(p Params, ret *int) error {
	*ret = p.Num1 + p.Num2
	return nil
}

// Diminish 求差服务
func (r *Response) Diminish(p Params, ret *int) error {

	if p.Num1 > p.Num2 {
		*ret = p.Num1 - p.Num2
	} else {
		*ret = p.Num2 - p.Num1
	}
	fmt.Println(*ret)

	return nil
}

// RPC 服务端三步走
func main() {
	var rep Response

	// 1. 注册服务
	err := rpc.Register(&rep)
	if err != nil {
		fmt.Println("服务注册失败！")
		return
	}

	// 2. 监听服务
	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("监听失败！")
		return
	}

	// 3. 循环获取客户端新连接
	for {
		accept, err := listen.Accept()
		if err != nil {
			continue
		}

		// 启动一个协程去处理
		go func() {
			fmt.Println("新客户端已连接！")
			jsonrpc.ServeConn(accept)
		}()
	}
}
