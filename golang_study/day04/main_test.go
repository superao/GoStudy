package day04

import (
	"reflect"
	"testing"
)

// Map 类型的声明和初始化
func TestMapInit(t *testing.T) {
	// Map 声明并初始化
	m := map[int]int{1: 1, 2: 4, 3: 6}
	//t.Log(m, len(m), cap(m))
	t.Log(m, len(m))

	// Map 声明操作
	m1 := map[string]int{}
	m1["str"] = 1
	m1["apple"] = 2
	m1["abc"] = 3
	t.Log(m1, len(m1))

	// 使用 make 初始化 Map 类型
	m2 := make(map[string]int, 10) // type cap，map初始化时没有个数，因为有些类型无法满足对应类型的零值
	m2["apple"] = 1                // 因此map类型初始化的时候无法指定元素个数
	m2["hello"] = 2
	m2["world"] = 3
	t.Log(m2, len(m2))
}

// Map类型变量的访问与设置
func TestMapGet(t *testing.T) {
	m := map[string]int{}
	// 访问map中不存在的值
	t.Log(m["str"])

	// 访问map中存在的值
	m["apple"] = 0
	t.Log(m["apple"])

	// 采用 map 访问时传递的第二个参数去辨别到底是零值还是值本身就那个样子
	if v, ok := m["str"]; !ok {
		t.Log("该值不存在！")
	} else {
		t.Log("该值为: ", v)
	}
}

// Map元素的遍历
func TestMapRange(t *testing.T) {
	// 声明并初始化一个 map 类型
	m := map[string]int{"str":1, "apple":1, "hello":2}

	// 遍历整个 map 元素
	for key, value := range m {
		t.Log(key, value)
	}
}

// Map实现工厂模式
func TestMapFactory(t *testing.T) {
	m := map[string]func(){}
	m["hello"] = func() {
		t.Log("hello")
	}
	m["world"] = func() {
		t.Log("world")
	}
	m["kuki"] = func() {
		t.Log("kuki")
	}

	// 向 map 传递什么，则就执行什么样的函数
	m["hello"]()
	m["world"]()
	m["kuki"]()

	// 泛型
	m1 := map[interface{}]interface{}{}
	m1["sum"] = func(){
		t.Log("sum")
	}

	t.Log(reflect.TypeOf(m1))   // map[interface {}]interface {}

	//m1["sum"]()  // 虽然可以赋值函数类型，但是无法通过接口类型去调用函数，这种是无法实现工厂模式的
}

// Map 实现 Set 操作
func TestMapSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true

	//m[1] = false

	if m[1] {
		t.Log("元素是存在的")
	}

	if !m[2] {
		t.Log("该元素是不存在的")
	}

	t.Log(len(m))
}




