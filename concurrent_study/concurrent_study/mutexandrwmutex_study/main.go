package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x     int = 0
	wg    sync.WaitGroup
	//mutex sync.Mutex
	rwmutex sync.RWMutex	// 读写分离的数量级越大，则读写锁的性能优势更高
)

// 对 x 进行写入操作
func write() {
	//mutex.Lock()
	rwmutex.Lock()
	x++
	time.Sleep(50 * time.Microsecond)
	//mutex.Unlock()
	rwmutex.Unlock()

	wg.Done()
}

// 对 x 进行读取操作
func read() {
	//mutex.Lock()
	rwmutex.RLock()
	//fmt.Println("获取的值为: ", x)
	time.Sleep(50 * time.Microsecond)
	//mutex.Unlock()
	rwmutex.RUnlock()

	wg.Done()
}

// 本示例演示使用互斥锁和读写锁来解决
func main() {
	// 记录开始时间
	now := time.Now()

	// 创建多个 goroutine 进行写入操作
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}

	// 创建多个 goroutine 进行读取操作(读写相差一个数量级)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}

	// 统计整个时间
	t := time.Since(now)
	fmt.Println("整个过程消耗的时间为: ", t)

	wg.Wait()
}
