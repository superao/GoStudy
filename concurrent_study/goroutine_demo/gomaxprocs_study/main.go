package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

func hello() {
	fmt.Println("hello ")
	wg.Done()
}

func world() {
	fmt.Println("world !")
	wg.Done()
}

// 本示例演示设置 内核态线程的数量 去执行 go 程序
func main() {
	//runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(8) // CPU 核心数越多，那么多个 goroutine 并行的概率就越多
	// 创建多个线程去执行 A 逻辑
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go hello()
	}

	// 创建多个线程去执行 B 逻辑
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go world()
	}

	wg.Wait()
}
