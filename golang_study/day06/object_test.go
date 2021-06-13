// 本示例演示 Go 语言如何实现面向对象
package day06

import (
	"fmt"
	"testing"
)
// 封装数据和行为
type Person struct {
	Id int64
	Name string
}

// 这种定义⽅式在实例对应⽅法被调⽤时，实例的成员会进⾏值复制，函数内的修改不会影响外面的值
func (p Person)Say() {
	fmt.Println("hello!")
	p.Name = "apple"
}

// 这种定义可以避免第一种方式带来的内存拷贝，并且函数内修改的内容，函数外是可见的
func (p *Person)Moving() {
	fmt.Println("moving")
	p.Id = 21
}

// 不管行为采用什么方式绑定，值类型的变量和指针类型的变量都是可以对其调用的
func TestPerson(t *testing.T) {
	p := Person{18, "str"}
	p.Say()
	p.Moving()

	p1 := &Person{21, "ao"}
	p1.Say()
	p1.Moving()
}

func TestPerson1(t *testing.T) {
	p := Person{18, "str"}
	p.Say()
	p.Moving()
	t.Log(p)  // 21 str
}

// 接口类型
type Pet interface {
	Say()
}

// 实现类型
type dog struct {

}

// 实现接口方法
func (d *dog)Say() {
	fmt.Println("wang!")
}

func TestInterface(t *testing.T) {
	var p Pet = new(dog)    // 接口类型的变量只接受指针类型
	p.Say()
}






