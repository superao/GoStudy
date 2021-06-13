package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc_study/message/message"
	"net"
)

// 定义服务端类型
type server struct {

}

// SayHello 定义服务端的服务
func (s *server)SayHello(ctx context.Context, person *message.ClientId) (*message.Person, error) {
	// 将接受过来的客户端 id 与 hello world 组合起来
	msg := message.Person{
		Id: person.Id,
		Name: "hello world!",
	}

	// 将服务处理完的消息返回给客户端
	return &msg, nil
}

func main() {
	// 1. 启动客户端服务，监听 9000 端口
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("服务启动失败~")
		return
	}

	// 2. 创建 grpc 服务器
	s := grpc.NewServer()

	// 3. 在 grpc 服务端注册当前的服务，即服务注册(register: 登记)
	message.RegisterSayHelloServer(s, &server{})

	// 4. 在 grpc 服务端注册服务器的反射服务，即向外暴漏出当前 grpc 服务器上所有已经注册过的服务
	reflection.Register(s)

	// 5. Serve 方法接受传入的连接，然后为每一个连接创建一个 ServerTransport 和 server 的 goroutine
	// 该 goroutine 服务 grpc 请求，然后调用已注册的处理程序来响应它们
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
