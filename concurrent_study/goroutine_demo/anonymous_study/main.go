package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

// 本示例演示在使用 goroutine 的时候采用匿名函数所导致的一些问题
func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		//go func() {
		//	fmt.Println("hello world, ", i)
		//	wg.Done()
		//}()
		go func(i int) {
			fmt.Println("hello world, ", i)
			wg.Done()
		}(i) // 将变量 i 传毒进去就不会出现函数闭包问题
	}

	wg.Wait()
}

// 输出结果:
//hello world,  5
//hello world,  5
//hello world,  5
//hello world,  5
//hello world,  5
//从结果中可以看出 goroutine 所执行的函数中并没有变量 i, 因此在使用 i 的时候必须在上一层作用域中去寻找
//此时有可能 i 已经变成了 5, goroutine 才获取到了它，因此结果不符合预期。
