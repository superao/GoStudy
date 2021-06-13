package main

import (
	"fmt"
	"sync"
)

var (
	x = 0
	wg sync.WaitGroup
)

// 累加操作
func add() {
	x++
	wg.Done()
}

// 本示例演示竞态条件所带来的影响，即不加锁操作全局变量 x
func main() {
	// 创建 100 个 goroutine 去进行累加操作
	for i := 0; i < 100; i++ {
		go add()
		wg.Add(1)
	}

	// 查看并发之后的结果是不是 100
	fmt.Println("并发后的结果为: ", x)		// 多次运行后结果都是不一致的，即竞态问题
	wg.Wait()
}
