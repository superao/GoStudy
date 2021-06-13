package main

import "fmt"

// 本示例演示使用 select 同时操作多个 channel 进行数据读写
func main() {
	// 创建三个 channel 类型的变量
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	ch3 := make(chan int, 10)

	// 向 ch1 中写入一个数据
	ch1 <- 10

	select {
		case value := <- ch1:
			fmt.Println("从 ch1 中接受到了数据: ", value)
		case ch2 <- 10:
			fmt.Println("向 ch2 中写入了一个数据~")
		default:
			ch3 <- 20
			fmt.Println("向 ch3 中写入了一个数据~")
	}
}
