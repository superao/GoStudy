package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup // 定义等待组，其作用类似于条件变量
)

func hello() {
	fmt.Println("hello ")
	wg.Done()
}

func world() {
	fmt.Println("world!")
	wg.Done()
}

// 本示例演示多个 goroutine 之间如何进行同步操作
func main() {
	// 启动多个 goroutine 去执行 hello 函数
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go hello()
	}

	// 启动多个 goroutine 去执行 world 函数
	for i := 1; i < 5; i++ {
		wg.Add(1)
		go world()
	}

	// 等待所有的 goroutine 执行完毕之后，再退出
	wg.Wait()
}
