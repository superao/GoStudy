package main

// 引入本地同一项目中的包，只需要从 src 目录下开始写写到具体的某个包就行
import "module_study/mypackage"     // 同一项目的包
import "mymodule/myhello"           // 不同项目的包

func main() {
	mypackage.Hello()
	myhello.World()
}
