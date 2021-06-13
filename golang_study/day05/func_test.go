package day05

import (
	"errors"
	"fmt"
	"testing"
	"time"
	"unsafe"
)

// 本示例学习函数相关的基础知识
// 演示函数是具有多个返回值
func myFunc() (string, int, string, error){
	return "hello", 888, "world", errors.New("error 255")
}

// 通常函数在使用的时候只返回两部分值，一部分是值，另一部分是错误信息
func myFunc1(op int) (string, error) {
	if op == 1 {
		return "hello world!", nil
	} else {
		return "", errors.New("出现了一个错误！")
	}
}


func TestFunction(t *testing.T) {
	t.Log(myFunc())

	t.Log(myFunc1(1))
	t.Log(myFunc1(10))
}

// 如果不指定传递的是指针类型，那么所有的都是值传递
func FuncByValue(op int) {
	fmt.Printf("这是数字是: %d, 地址为: %x\n", op, &op)
}

func FunByMap(m map[string]int) {
	fmt.Printf("map 的地址为: %x", unsafe.Pointer(&m))
}

func TestFuncByValue(t *testing.T) {
	num := 10
	t.Logf("地址为: %x\n", &num)   // 经过测试发现，两个变量的地址并不相同，因此可以
	FuncByValue(num)

	var m map[string]int = map[string]int{"hello":1, "world":2}
	t.Logf("地址为: %x", unsafe.Pointer(&m))    // 经过测试发现，两个 map 的地址依然是不相同的
	FunByMap(m)
}


// 可变参函数的使用
func funcMutil(op ...int) int {
	ret := 0
	for _, e := range op {
		ret += e
	}
	return ret
}

func TestMutilFunc(t *testing.T) {
	t.Log(funcMutil(1,2,3,4,5))
	t.Log(funcMutil(1,2,3,4))
}

// 将函数作为多种用途：变量的值、函数参数、函数返回值
// 在原有函数的基础上封装上一层时间统计耗时的功能
func timeFunc(f func(op int)int) func(int)int {
	tmpFunc := func(op int) int {
		t := time.Now()
		ret := f(op)
		fmt.Printf("函数耗时统计结果为: %f", time.Since(t).Seconds())
		return ret
	}

	return tmpFunc
}

func testSleep(op int) int {
	time.Sleep(time.Second * 2)
	fmt.Println("hello world!")

	return op
}

func TestFuncTime(t *testing.T) {
	timeF:= timeFunc(testSleep)

	t.Log(timeF(10))
}

// defer 函数的使用
func funcDefer() {
	defer func() {
		fmt.Println("defer function!")    // 延迟调用
	}()
	fmt.Println("hello")
}

func TestDefer(t *testing.T) {
	funcDefer()
}

func recoverPanic() {
	defer func() {
		if err := recover(); err == "200" {
			fmt.Println("200 recover!")
		}
	}()
	fmt.Println("hello world!")
	panic("200")
}

func TestRecoverPanic(t *testing.T) {
	recoverPanic()
}


