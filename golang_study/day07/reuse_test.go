// 本文件演示 Go 语言中复用操作
package day07

import (
	"errors"
	"fmt"
	"testing"
)

//type father struct {
//	Id int64
//	Name string
//}
//
//func (f *father)Say() {
//	fmt.Println("hello~")
//}
//
//type son struct {
//	f father
//	Age int32
//}
//
//func (s *son)Say() {
//	fmt.Println("hi~")
//}
//
//func TestReuse(t *testing.T) {
//	s := son{}
//	s.f.Id = 123456
//	s.f.Say()   // hello~
//	s.Age = 21
//	s.Say()		// hi~
//}

type father struct {
	Id int64
	Name string
}

func (f *father)Say() {
	fmt.Println("hello~")
}

type son struct {
	father     // 匿名类型嵌入，可直接使用其数据和方法，特殊语法，不是继承
}

func TestReuse(t *testing.T) {
	s := son{}
	s.Name = "ao"
	s.Say()  // hello~

	//var f father = son{}          // 编译错误
	//var f1 *father = &son{}       // 编译错误
	//var f2 father = father(son{}) // 编译错误
}

// 空接口与断言
func Interface(op interface{}) {
	if v, ok := op.(int); ok {
		fmt.Println("int ", v)
	}
	if v, ok := op.(string); ok {
		fmt.Println("string ", v)
	}
}
func InterfaceBySwitch(op interface{}) {
	switch v := op.(type) {
	case int:
		fmt.Println("int ", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("unknown!")
	}
}
func TestInterface(t *testing.T) {
	Interface(10)
	Interface("hello!")
	InterfaceBySwitch(10)
	InterfaceBySwitch(false)
}

// 错误检测
func testNumber(n int) error {
	if n > 10 {
		return errors.New("大了！")
	}
	if n < 0 {
		return errors.New("小了！")
	}

	fmt.Println("恭喜你！")
	return nil
}

func TestError(t *testing.T) {
	err := testNumber(12)
	if err != nil {
		if err.Error() == "大了！" {
			t.Log("猜大了~")
		}
		if err.Error() == "小了！" {
			t.Log("猜小了~")
		}
	}
}

func TestErrorType(t *testing.T) {
	err1 := errors.New("error")
	err2 := errors.New("error")
	if err1 != err2 {
		t.Log("不相同~")    // 不相同~
	}

	t.Logf("err1 type: %T, err2 type: %T", err1, err2)
	// err1 type: *errors.errorString, err2 type: *errors.errorString
	// 因此两个 error 变量没有办法直接进行比较

	if err1.Error() == err2.Error() {
		t.Log("相同~")     // 相同~
	}
}




