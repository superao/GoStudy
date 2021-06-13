package main

import "fmt"

// 从 in 通道中读取数据，计算之后写入到 out 通道中
// 单向通道在区分的时候以 chan 关键字为核心去进行区分
func function(out chan<- int, in <-chan int) {
	//in <- 10 编译不通过
	// 从 channel 中读取数据
	for value := range in {
		v := value + 100

		//tmp := <- out   编译不通过
		// 将计算结果写入到 channel 中
		out <- v
	}
	close(out)
}

// 本示例演示单向通道的作用
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100) // 注意: 缓冲区大小一定要够，否则就会阻塞住

	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)

	go function(ch2, ch1)

	for value := range ch2 {
		fmt.Println(value)
	}
}
