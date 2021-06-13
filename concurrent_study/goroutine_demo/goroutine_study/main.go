package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello kukizhang!")
}

// 学习使用 goroutine 的例子
func main() {
	//hello()   // main 函数默认创建的 goroutine 会按照先后顺序执行两句打印代码
	go hello() // 当启动一个新的 goroutine 的时候，两句打印代码的执行没有先后顺序，它们是同时执行的
	fmt.Println("hello world!")
	time.Sleep(1 * time.Second) // Sleep 函数默认睡眠 1 纳秒
}
