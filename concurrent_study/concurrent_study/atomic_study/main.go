package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 全局变量 x
var x int32 = 0
var wg sync.WaitGroup

func add(){
	atomic.AddInt32(&x, 1)	 // 原子操作，性能更高，并且不会出现并发问题
	wg.Done()
}

func get(){
	atomic.LoadInt32(&x)
	fmt.Println("当前 x 的值为: ", x)
	wg.Done()
}

// 本示例演示 sync 包中提供的原子操作
func main() {
	// 创建多个 goroutine 去累加值
	for i:=0; i < 1000; i++ {
		go add()
		wg.Add(1)
	}

	wg.Wait()

	// 查看累加后的值
	fmt.Println("当前 x 的值为: ", x)
}

