package main

import "fmt"

// 产生 1 ~ 100 之间的数，然后写入在 channel 中
func input(in chan int) {
	// 生产 1 ~ 100 之间的数据
	for i:=1; i <= 100; i++ {
		// 向 channel 中写入数据
		in <- i
	}

	// 数据传输完毕之后，关闭 channel 类型
	close(in)
}

// 从 channel 中读取数据，并且将每一个数据进行平方后，再写入到另一个 channel 中
func compute(out chan int, in chan int) {
	// 循环从 channel 中读取数据
	for value := range in {
		// 平方
		v := value * value
		// 写入到 channel 中
		out <- v
	}

	// 写入完毕之后，关闭 channel 类型
	close(out)
}

// 本示例演示 channel 在 goroutine 之间的通信
func main() {
	// 初始化 channel 类型
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go input(ch1)
	go compute(ch2, ch1)

	// 从 channel 中读取数据并输出
	for {
		value, ok := <- ch2
		if !ok {
			break
		}

		fmt.Println(value)
	}
}
