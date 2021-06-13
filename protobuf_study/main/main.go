package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protobuf_study/person"
)

func main() {
	// 定义 person 变量
	p := person.Person{
		Name: "zhangsan",
		Id: 123456,
		Sex: person.Gender_boy,
	}

	// 输出该变量的所有字段
	// 这里取地址符和不取地址符是不相同的，具体得了解原理后才清楚
	fmt.Println(p)

	// 将该变量序列化成二进制格式输出
	data, err := proto.Marshal(&p)
	if err != nil {
		fmt.Println("proto 序列化错误，", err)
		return
	}
	fmt.Println("序列化后的数据为：", data)

	// 将序列化后的二进制数据进行后序列化还原成原数据格式
	var per person.Person
	err = proto.Unmarshal(data, &per)
	if err != nil {
		fmt.Println("proto 反序列化错误，", err)
		return
	}
	fmt.Println(&per)

	// 获取该变量的每个字段
	fmt.Println(per.GetName())
	fmt.Println(per.GetId())
	fmt.Println(per.Sex)
}
