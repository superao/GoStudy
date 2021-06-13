package main

import (
	"context"
	"fmt"
	"grpc_study/message/message"

	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := message.NewSayHelloClient(conn)
	// 调用服务端的SayHello
	r, err := c.SayHello(context.Background(), &message.ClientId{Id: 123456})
	if err != nil {
		fmt.Printf("could not greet: %v", err)
	}
	fmt.Println(r.GetId(), r.GetName())
}
