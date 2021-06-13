package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func compute(i int, in chan int) {
	// 每一个协程都是循环的从通道任务队列中拿数据，然后进行处理
	// channel 本身是线程安全的，多个 goroutine 在对同一个 channel 进行读写的时候是不需要要进行加锁的！！！
	for task := range in {
		time.Sleep(50 * time.Microsecond)
		fmt.Println("当前的协程ID为: ", i, " 任务序号为: ", task, "处理完毕！")
	}
	wg.Done()
}

// 本示例演示 go 语言中的 goroutine 协程池
func main() {
	// 创建任务的缓冲队列(通道)
	ch1 := make(chan int, 100)

	// 创建消费者协程池
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go compute(i, ch1)
	}

	// 创建 100 个任务
	for i := 1; i<=100000; i++ {
		ch1 <- i
	}
	close(ch1)

	wg.Wait()
}
