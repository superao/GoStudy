// Package series 本示例演示包的可复用性
package series

import "fmt"

func init() {
	fmt.Println("init func 1!")
}

func init() {
	fmt.Println("init func 2!")
}

func Series(num int) {
	ret := []int{1, 1}
	for i := 2; i < num; i++ {
		ret = append(ret, ret[i-2] + ret[i-1])
	}

	fmt.Printf("结果为: %v", ret)
}
