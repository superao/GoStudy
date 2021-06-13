package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
	m map[int]int	// 当 goroutine 数量较多时，就会出现运行时错误
	safem sync.Map
)

// 初始化所有的全部变量，自动被调用
func init() {
	fmt.Println("初始化 map 变量开始")
	m = make(map[int]int, 100)
	fmt.Println("初始化 map 变量结束")
}

// 操作 map 的函数
func set(key int, value int) {
	m[key] = value
	wg.Done()
}

func get(key int) {
	fmt.Println("获取到的值为: ", m[key])
	wg.Done()
}

// 操作 sync.map 的函数
func safemSet(key int, value int) {
	safem.Store(key, value)
	wg.Done()
	return
}

func safemGet(key int) {
	value, ok := safem.Load(key)
	if !ok {
		fmt.Println("获取数据失败！")
		return
	}

	fmt.Println("获取的值为: ", value)
	wg.Done()
	return
}

// 本示例演示并发不安全的 map 和 并发安全的 map 操作
func main() {
	// 创建多个 goroutine 去操作同一个 map
	for i := 0; i < 1000; i++ {
		//go set(i, i + 100)
		go safemSet(i, i + 1000)
		wg.Add(1)
	}

	// 创建多个 goroutine 去读取同一个 map
	for i := 0; i < 1000; i++ {
		//go get(i)
		go safemGet(i)
		wg.Add(1)
	}

	wg.Wait()
}
