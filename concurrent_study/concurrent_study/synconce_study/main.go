package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton 定义对象类型
type Singleton struct{}

func (s *Singleton) hello() {
	fmt.Println("hello world!")
}

// 定义单例对象
var singleton *Singleton
var once sync.Once

// 定义创建单例对象的全局函数
func create() {
	once.Do(func() {
		singleton = &Singleton{}
	})
	fmt.Printf("对象的地址为: %p\n", singleton)

	return
}

// 本示例演示采用 sync.once 提供的方法去实现单例模式
func main() {
	// 创建多个 goroutine 去创建单例对象，观察获取的是否是同一个对象
	for i := 0; i < 50; i++ {
		go create()		// 多个 goroutine 拿到的对象地址是相同的，不会出现并发不安全的问题出现
	}

	time.Sleep(3 * time.Second)
}
