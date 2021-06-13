package main

import (
	"fmt"
	"time"
)

func recv(ch chan int)  {
	// 从 channel 中读取数据
	value := <- ch
	fmt.Println(value)
}

// 本示例演示如何创建一个可用的 channel，并使用它进行通信
func main() {
	// 1. 声明 channel 类型
	// 声明一个 channel 类型，ch 为 nil
	//var ch chan int
	//fmt.Printf("chan 变量: %v", ch)
	// 声明一个 channel 类型，并进行初始化

	// 2. 声明 channel 类型，并使用
	//var ch chan int
	//ch = make(chan int, 10)
	//fmt.Printf("chan 变量: %v\n", ch)
	//
	//// 向 channel 中写入数据
	//ch <- 10
	//
	//// 从 channel 中读取数据，并操作
	//value := <- ch
	//fmt.Println(value)

	// 3. 无缓冲区的 channel 类型
	//ch := make(chan int)
	//ch <- 10	// 没有人接受，channel 阻塞住，直接导致死锁

	// 4. 无缓冲区的 channel 类型，并有 goroutine 接受
	ch := make(chan int)
	go recv(ch)  // 这里必须写在前面，否则就会发生死锁
	// 向 channel 中写入数据
	ch <- 10
	close(ch)
	time.Sleep(1 * time.Second)
}
